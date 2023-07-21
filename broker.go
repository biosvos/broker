package broker

import "context"

type Publisher interface {
	Publish(ctx context.Context, message []byte) error
}

type Subscriber interface {
	Subscribe(ctx context.Context) ([]byte, error)
}

type Broker interface {
	NewPublisher(topic string) (Publisher, error)
	NewSubscriber(topic string) (Subscriber, error)
	Close()
}
