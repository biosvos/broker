package rabbitmq

import (
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

func (b *Broker) NewPublisher(topic string) (broker.Publisher, error) {
	err := b.channel.ExchangeDeclare(topic, "direct", false, true, false, false, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return NewPublisher(b.channel, topic), nil
}

func (b *Broker) NewSubscriber(topic string) (broker.Subscriber, error) {
	err := b.channel.ExchangeDeclare(topic, "direct", false, true, false, false, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	queue, err := b.channel.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = b.channel.QueueBind(queue.Name, "", topic, false, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	consume, err := b.channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return NewSubscriber(consume), nil
}

func (b *Broker) Close() {
	err := b.channel.Close()
	if err != nil {
		log.Println(err)
	}
}
