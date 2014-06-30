package nsq

import (
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// Producer is a high-level type to publish to NSQ.
//
// A Producer instance is 1:1 with a destination `nsqd`
// and will lazily connect to that instance (and re-connect)
// when Publish commands are executed.
type Producer struct {
	id     int64
	addr   string
	conn   *Conn
	config Config

	logger *log.Logger
	logLvl LogLevel

	responseChan  chan []byte
	errorChan     chan []byte
	ioErrorChan   chan error
	heartbeatChan chan int
	closeChan     chan int

	transactionChan chan *ProducerTransaction
	transactions    []*ProducerTransaction
	state           int32

	concurrentProducers int32
	stopFlag            int32
	exitChan            chan int
	wg                  sync.WaitGroup
	guard               sync.Mutex
}

// ProducerTransaction is returned by the async publish methods
// to retrieve metadata about the command after the
// response is received.
type ProducerTransaction struct {
	cmd      *Command
	doneChan chan *ProducerTransaction
	Error    error         // the error (or nil) of the publish command
	Args     []interface{} // the slice of variadic arguments passed to PublishAsync or MultiPublishAsync
}

func (t *ProducerTransaction) finish() {
	if t.doneChan != nil {
		t.doneChan <- t
	}
}

// NewProducer returns an instance of Producer for the specified address
//
// The only valid way to create a Config is via NewConfig, using a struct literal will panic.
// After Config is passed into NewProducer the values are no longer mutable (they are copied).
func NewProducer(addr string, config *Config) (*Producer, error) {
	config.assertInitialized()
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	p := &Producer{
		id: atomic.AddInt64(&instCount, 1),

		addr:   addr,
		config: *config,

		logger: log.New(os.Stderr, "", log.Flags()),
		logLvl: LogLevelInfo,

		transactionChan: make(chan *ProducerTransaction),
		exitChan:        make(chan int),
		responseChan:    make(chan []byte),
		errorChan:       make(chan []byte),
		ioErrorChan:     make(chan error),
		heartbeatChan:   make(chan int),
		closeChan:       make(chan int),
	}
	return p, nil
}

// SetLogger assigns the logger to use as well as a level
func (w *Producer) SetLogger(logger *log.Logger, lvl LogLevel) {
	w.logger = logger
	w.logLvl = lvl
}

// String returns the address of the Producer
func (w *Producer) String() string {
	return w.addr
}

// Stop initiates a graceful stop of the Producer (permanent)
//
// NOTE: this blocks until completion
func (w *Producer) Stop() {
	w.guard.Lock()
	if !atomic.CompareAndSwapInt32(&w.stopFlag, 0, 1) {
		w.guard.Unlock()
		return
	}
	w.log(LogLevelInfo, "stopping")
	close(w.exitChan)
	w.close()
	w.guard.Unlock()
	w.wg.Wait()
}

// PublishAsync publishes a message body to the specified topic
// but does not wait for the response from `nsqd`.
//
// When the Producer eventually receives the response from `nsqd`,
// the supplied `doneChan` (if specified)
// will receive a `ProducerTransaction` instance with the supplied variadic arguments
// and the response error if present
func (w *Producer) PublishAsync(topic string, body []byte, doneChan chan *ProducerTransaction,
	args ...interface{}) error {
	return w.sendCommandAsync(Publish(topic, body), doneChan, args)
}

// MultiPublishAsync publishes a slice of message bodies to the specified topic
// but does not wait for the response from `nsqd`.
//
// When the Producer eventually receives the response from `nsqd`,
// the supplied `doneChan` (if specified)
// will receive a `ProducerTransaction` instance with the supplied variadic arguments
// and the response error if present
func (w *Producer) MultiPublishAsync(topic string, body [][]byte, doneChan chan *ProducerTransaction,
	args ...interface{}) error {
	cmd, err := MultiPublish(topic, body)
	if err != nil {
		return err
	}
	return w.sendCommandAsync(cmd, doneChan, args)
}

// Publish synchronously publishes a message body to the specified topic, returning
// the an error if publish failed
func (w *Producer) Publish(topic string, body []byte) error {
	return w.sendCommand(Publish(topic, body))
}

// MultiPublish synchronously publishes a slice of message bodies to the specified topic, returning
// the an error if publish failed
func (w *Producer) MultiPublish(topic string, body [][]byte) error {
	cmd, err := MultiPublish(topic, body)
	if err != nil {
		return err
	}
	return w.sendCommand(cmd)
}

func (w *Producer) sendCommand(cmd *Command) error {
	doneChan := make(chan *ProducerTransaction)
	err := w.sendCommandAsync(cmd, doneChan, nil)
	if err != nil {
		close(doneChan)
		return err
	}
	t := <-doneChan
	return t.Error
}

func (w *Producer) sendCommandAsync(cmd *Command, doneChan chan *ProducerTransaction,
	args []interface{}) error {
	// keep track of how many outstanding producers we're dealing with
	// in order to later ensure that we clean them all up...
	atomic.AddInt32(&w.concurrentProducers, 1)
	defer atomic.AddInt32(&w.concurrentProducers, -1)

	if atomic.LoadInt32(&w.state) != StateConnected {
		err := w.connect()
		if err != nil {
			return err
		}
	}

	t := &ProducerTransaction{
		cmd:      cmd,
		doneChan: doneChan,
		Args:     args,
	}

	select {
	case w.transactionChan <- t:
	case <-w.exitChan:
		return ErrStopped
	}

	return nil
}

func (w *Producer) connect() error {
	w.guard.Lock()
	defer w.guard.Unlock()

	if atomic.LoadInt32(&w.stopFlag) == 1 {
		return ErrStopped
	}

	if !atomic.CompareAndSwapInt32(&w.state, StateInit, StateConnected) {
		return ErrNotConnected
	}

	w.log(LogLevelInfo, "(%s) connecting to nsqd", w.addr)

	conn := NewConn(w.addr, &w.config)
	conn.SetLogger(w.logger, w.logLvl, fmt.Sprintf("%3d (%%s)", w.id))
	conn.Delegate = &producerConnDelegate{w}

	_, err := conn.Connect()
	if err != nil {
		conn.Close()
		w.log(LogLevelError, "(%s) error connecting to nsqd - %s", w.addr, err)
		atomic.StoreInt32(&w.state, StateInit)
		return err
	}
	w.conn = conn

	w.wg.Add(1)
	go w.router()

	return nil
}

func (w *Producer) close() {
	if !atomic.CompareAndSwapInt32(&w.state, StateConnected, StateDisconnected) {
		return
	}
	w.conn.Close()
	go func() {
		// we need to handle this in a goroutine so we don't
		// block the caller from making progress
		w.wg.Wait()
		atomic.StoreInt32(&w.state, StateInit)
	}()
}

func (w *Producer) router() {
	for {
		select {
		case t := <-w.transactionChan:
			w.transactions = append(w.transactions, t)
			err := w.conn.WriteCommand(t.cmd)
			if err != nil {
				w.log(LogLevelError, "(%s) sending command - %s", w.conn.String(), err)
				w.close()
			}
		case data := <-w.responseChan:
			w.popTransaction(FrameTypeResponse, data)
		case data := <-w.errorChan:
			w.popTransaction(FrameTypeError, data)
		case <-w.heartbeatChan:
		case <-w.ioErrorChan:
			w.close()
		case <-w.closeChan:
			goto exit
		case <-w.exitChan:
			goto exit
		}
	}

exit:
	w.transactionCleanup()
	w.wg.Done()
	w.log(LogLevelInfo, "exiting router")
}

func (w *Producer) popTransaction(frameType int32, data []byte) {
	t := w.transactions[0]
	w.transactions = w.transactions[1:]
	if frameType == FrameTypeError {
		t.Error = ErrProtocol{string(data)}
	}
	t.finish()
}

func (w *Producer) transactionCleanup() {
	// clean up transactions we can easily account for
	for _, t := range w.transactions {
		t.Error = ErrNotConnected
		t.finish()
	}
	w.transactions = w.transactions[:0]

	// spin and free up any writes that might have raced
	// with the cleanup process (blocked on writing
	// to transactionChan)
	for {
		select {
		case t := <-w.transactionChan:
			t.Error = ErrNotConnected
			t.finish()
		default:
			// keep spinning until there are 0 concurrent producers
			if atomic.LoadInt32(&w.concurrentProducers) == 0 {
				return
			}
			// give the runtime a chance to schedule other racing goroutines
			time.Sleep(5 * time.Millisecond)
			continue
		}
	}
}

func (w *Producer) log(lvl LogLevel, line string, args ...interface{}) {
	var prefix string

	if w.logger == nil {
		return
	}

	if w.logLvl > lvl {
		return
	}

	switch lvl {
	case LogLevelDebug:
		prefix = "DBG"
	case LogLevelInfo:
		prefix = "INF"
	case LogLevelWarning:
		prefix = "WRN"
	case LogLevelError:
		prefix = "ERR"
	}

	w.logger.Printf("%-4s %3d %s", prefix, w.id, fmt.Sprintf(line, args...))
}

func (w *Producer) onConnResponse(c *Conn, data []byte) { w.responseChan <- data }
func (w *Producer) onConnError(c *Conn, data []byte)    { w.errorChan <- data }
func (w *Producer) onConnHeartbeat(c *Conn)             { w.heartbeatChan <- 1 }
func (w *Producer) onConnIOError(c *Conn, err error)    { w.ioErrorChan <- err }
func (w *Producer) onConnClose(c *Conn)                 { w.closeChan <- 1 }
