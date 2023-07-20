package rabbitmq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPublish(t *testing.T) {
	broker, err := NewBroker("guest", "guest", "127.0.0.1", 5672)
	require.NoError(t, err)

	err = broker.Publish("hi", []byte("hi"))
	require.NoError(t, err)

	message, err := broker.Subscribe("hi")
	require.NoError(t, err)
	require.EqualValues(t, "hi", message)
}
