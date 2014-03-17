package consumer

// Shortcuts for apps using only one consumer
var DefaultConsumer = NewConsumer()

// Look for Consumer.Connect() (on default consumer)
func Connect(addr string) error {
	return DefaultConsumer.Connect(addr)
}

// Look for Consumer.ConnectList() (on default consumer)
func ConnectList(addrs []string) error {
	return DefaultConsumer.ConnectList(addrs)
}

// Look for Consumer.ConnectLookupd() (on default consumer)
func ConnectLookupd(addr string) error {
	return DefaultConsumer.ConnectLookupd(addr)
}

// Look for Consumer.ConnectLookupdList() (on default consumer)
func ConnectLookupdList(addrs []string) error {
	return DefaultConsumer.ConnectLookupdList(addrs)
}

// Look for Consumer.Register() (on default consumer)
func Register(topic, channel string, maxInFlight int, handler func(*Message)) error {
	return DefaultConsumer.Register(topic, channel, maxInFlight, handler)
}

// Look for Consumer.Start() (on default consumer)
func Start(debug bool) {
	DefaultConsumer.Start(debug)
}
