package memory

import (
	"context"
	"errors"
	"github.com/biosvos/broker"
)

var _ broker.Subscriber = &Subscriber{}

func NewSubscriber(ch <-chan *Message) *Subscriber {
	return &Subscriber{
		ch: ch,
	}
}

type Subscriber struct {
	ch <-chan *Message
}

func (s *Subscriber) Subscribe(ctx context.Context) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("done context")
	case delivery := <-s.ch:
		return delivery.contents, nil
	}
}
