package rabbitmq

import (
	"context"
	"fmt"
	"github.com/biosvos/broker"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var _ broker.Broker = &Broker{}

func NewBroker(username, password, address string, port uint16) (*Broker, error) {
	dial, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v", username, password, address, port))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	channel, err := dial.Channel()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Broker{
		channel:    channel,
		subscriber: map[string]<-chan amqp.Delivery{},
	}, nil
}

type Broker struct {
	channel    *amqp.Channel
	subscriber map[string]<-chan amqp.Delivery
}

func (b *Broker) Publish(topic string, message []byte) error {
	err := b.channel.ExchangeDeclare(topic, "fanout", false, true, false, false, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	queue, err := b.channel.QueueDeclare(topic, false, false, false, false, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	err = b.channel.QueueBind(queue.Name, "", topic, false, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	err = b.channel.PublishWithContext(context.Background(), topic, "", false, false, amqp.Publishing{
		Body: message,
	})

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (b *Broker) Subscribe(topic string) ([]byte, error) {
	subscriber, err := b.claimSubscriber(topic)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	msg := <-subscriber
	return msg.Body, nil
}

func (b *Broker) claimSubscriber(topic string) (<-chan amqp.Delivery, error) {
	subscriber, ok := b.subscriber[topic]
	if ok {
		return subscriber, nil
	}
	consume, err := b.channel.Consume(topic, "", true, false, false, false, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	b.subscriber[topic] = consume
	return consume, nil
}

func (b *Broker) Close() {
	err := b.channel.Close()
	if err != nil {
		log.Println(err)
	}
}
