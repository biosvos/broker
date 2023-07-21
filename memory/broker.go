package memory

import (
	"github.com/biosvos/broker"
)

var _ broker.Broker = &Broker{}

func NewBroker() (*Broker, error) {
	b := &Broker{
		publisherChannel: make(chan *Message),
		topicChannels:    map[string][]chan *Message{},
		stopCh:           make(chan struct{}),
	}
	go b.process()
	return b, nil
}

type Broker struct {
	publisherChannel chan *Message
	topicChannels    map[string][]chan *Message
	stopCh           chan struct{}
}

func (b *Broker) process() {
	for {
		select {
		case msg := <-b.publisherChannel:
			for _, ch := range b.topicChannels[msg.topic] {
				ch <- msg
			}
		case <-b.stopCh:
			return
		}
	}
}

func (b *Broker) NewPublisher(topic string) (broker.Publisher, error) {
	return NewPublisher(topic, b.publisherChannel), nil
}

func (b *Broker) NewSubscriber(topic string) (broker.Subscriber, error) {
	ch := make(chan *Message)
	b.topicChannels[topic] = append(b.topicChannels[topic], ch)
	return NewSubscriber(ch), nil
}

func (b *Broker) Close() {
	b.stopCh <- struct{}{}
}
