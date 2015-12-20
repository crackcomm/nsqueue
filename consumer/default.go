package consumer

// DefaultConsumer - Shortcuts for apps using only one consumer.
var DefaultConsumer *Consumer

func init() {
	DefaultConsumer = New()
}

// Connect - Look for Consumer.Connect() (on default consumer)
func Connect(addrs ...string) error {
	return DefaultConsumer.Connect(addrs...)
}

// ConnectLookupd - Look for Consumer.ConnectLookupd() (on default consumer)
func ConnectLookupd(addrs ...string) error {
	return DefaultConsumer.ConnectLookupd(addrs...)
}

// Register - Look for Consumer.Register() (on default consumer)
func Register(topic, channel string, maxInFlight int, fnc Handler) error {
	return DefaultConsumer.Register(topic, channel, maxInFlight, fnc)
}

// Start - Look for Consumer.Start() (on default consumer)
func Start(debug bool) {
	DefaultConsumer.Start(debug)
}
