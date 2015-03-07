package consumer

// Shortcuts for apps using only one consumer
var DefaultConsumer *Consumer

func init() {
	DefaultConsumer = New()
}

// Connect - Look for Consumer.Connect() (on default consumer)
func Connect(addr string) error {
	return DefaultConsumer.Connect(addr)
}

// ConnectList - Look for Consumer.ConnectList() (on default consumer)
func ConnectList(addrs []string) error {
	return DefaultConsumer.ConnectList(addrs)
}

// ConnectLookupd - Look for Consumer.ConnectLookupd() (on default consumer)
func ConnectLookupd(addr string) error {
	return DefaultConsumer.ConnectLookupd(addr)
}

// ConnectLookupdList - Look for Consumer.ConnectLookupdList() (on default consumer)
func ConnectLookupdList(addrs []string) error {
	return DefaultConsumer.ConnectLookupdList(addrs)
}

// Register - Look for Consumer.Register() (on default consumer)
func Register(topic, channel string, maxInFlight int, fnc Handler) error {
	return DefaultConsumer.Register(topic, channel, maxInFlight, fnc)
}

// Start - Look for Consumer.Start() (on default consumer)
func Start(debug bool) {
	DefaultConsumer.Start(debug)
}
