package rabbitmq

import (
	"context"
	"github.com/biosvos/broker"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

var _ broker.Publisher = &Publisher{}

func NewPublisher(channel *amqp.Channel, topic string) *Publisher {
	return &Publisher{
		channel: channel,
		topic:   topic,
	}
}

type Publisher struct {
	channel *amqp.Channel
	topic   string
}

func (p *Publisher) Publish(ctx context.Context, message []byte) error {
	err := p.channel.PublishWithContext(ctx, p.topic, "", false, false, amqp.Publishing{
		Body: message,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
