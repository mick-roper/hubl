package common

import "errors"

type (
	// Message that comes from a provider
	Message struct {
		Topic string
		Data  []byte
	}

	// MessageProcessor wraps multiple providers
	MessageProcessor struct {
		dataChannel       chan Message
		providers         []MessageProvider
		subscriptionStore SubscriptionStore
	}

	// MessageProvider listens for messages
	MessageProvider interface {
		OnDataReceived(chan<- Message)
	}
)

// NewMessageProcessor creates a new instance of the message service
func NewMessageProcessor(subscriptionStore SubscriptionStore) (*MessageProcessor, error) {
	if subscriptionStore == nil {
		return nil, errors.New("subscriptionStore is nil")
	}

	return &MessageProcessor{
		dataChannel:       make(chan Message),
		providers:         []MessageProvider{},
		subscriptionStore: subscriptionStore,
	}, nil
}
