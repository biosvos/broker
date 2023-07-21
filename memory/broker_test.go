package memory

import (
	"context"
	"github.com/biosvos/broker"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBroker(t *testing.T) {
	brk, err := NewBroker()
	require.NoError(t, err)
	defer brk.Close()
	publisher, err := brk.NewPublisher("test")
	require.NoError(t, err)

	var subscribers []broker.Subscriber
	for i := 0; i < 3; i++ {
		subscriber, err := brk.NewSubscriber("test")
		require.NoError(t, err)
		subscribers = append(subscribers, subscriber)
	}

	err = publisher.Publish(context.Background(), []byte("hi"))
	require.NoError(t, err)

	for _, subscriber := range subscribers {
		msg, err := subscriber.Subscribe(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, "hi", msg)
	}
}
