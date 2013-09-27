package consumer

var DefaultConsumer = NewConsumer()

func Register(topic, channel string, maxInFlight int, handler func(*Message)) {
	DefaultConsumer.Register(topic, channel, maxInFlight, handler)
}

func ConnectToLookupd(addr string) error {
	return DefaultConsumer.ConnectToLookupd(addr)
}

func ConnectToNSQ(addr string) error {
	return DefaultConsumer.ConnectToNSQ(addr)
}

func ConnectToLookupds(addrs []string) error {
	return DefaultConsumer.ConnectToLookupds(addrs)
}

func ConnectToNSQs(addrs []string) error {
	return DefaultConsumer.ConnectToNSQs(addrs)
}

func Start(debug bool) {
	DefaultConsumer.Start(debug)
}
