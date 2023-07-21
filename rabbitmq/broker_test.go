package rabbitmq

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestPublish(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	broker, err := NewBroker("guest", "guest", "127.0.0.1", 5672)
	require.NoError(t, err)
	defer broker.Close()

	publisher, err := broker.NewPublisher("test")
	require.NoError(t, err)

	subscriber, err := broker.NewSubscriber("test")
	require.NoError(t, err)

	err = publisher.Publish(context.Background(), []byte("hi"))
	require.NoError(t, err)

	delivery, err := subscriber.Subscribe(context.Background())
	require.NoError(t, err)

	t.Log(string(delivery))
}
