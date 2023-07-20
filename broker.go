package broker

type Broker interface {
	Publish(topic string, message []byte) error
	Subscribe(topic string) ([]byte, error)

	Close()
}
