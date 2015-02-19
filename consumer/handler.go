package consumer

// Handler - Function that handles incoming message.
type Handler func(*Message)

// AsyncHandler - Executes handler in goroutine.
func AsyncHandler(h Handler) Handler {
	return func(msg *Message) {
		go h(msg)
	}
}
