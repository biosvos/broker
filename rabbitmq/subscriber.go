package rabbitmq

import (
	"context"
	"github.com/biosvos/broker"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

var _ broker.Subscriber = &Subscriber{}

func NewSubscriber(ch <-chan amqp.Delivery) *Subscriber {
	return &Subscriber{
		ch: ch,
	}
}

type Subscriber struct {
	ch <-chan amqp.Delivery
}

func (s *Subscriber) Subscribe(ctx context.Context) ([]byte, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	delivery := <-s.ch
	return delivery.Body, nil
}
