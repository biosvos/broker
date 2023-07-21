package memory

import (
	"context"
	"github.com/biosvos/broker"
	"github.com/pkg/errors"
)

var _ broker.Publisher = &Publisher{}

func NewPublisher(topic string, ch chan *Message) *Publisher {
	return &Publisher{
		topic: topic,
		ch:    ch,
	}
}

type Publisher struct {
	topic string
	ch    chan *Message
}

func (p *Publisher) Publish(ctx context.Context, message []byte) error {
	err := ctx.Err()
	if err != nil {
		return errors.WithStack(err)
	}
	p.ch <- &Message{
		topic:    p.topic,
		contents: message,
	}
	return nil
}
