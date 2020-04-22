package common

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

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
		client            http.Client
		logger            *log.Logger
		stopped           bool
	}

	// MessageProvider listens for messages
	MessageProvider interface {
		OnDataReceived(chan<- Message)
	}
)

// NewMessageProcessor creates a new instance of the message service
func NewMessageProcessor(subscriptionStore SubscriptionStore, logger *log.Logger) (*MessageProcessor, error) {
	if subscriptionStore == nil {
		return nil, errors.New("subscriptionStore is nil")
	}

	return &MessageProcessor{
		dataChannel:       make(chan Message),
		providers:         []MessageProvider{},
		subscriptionStore: subscriptionStore,
		client:            http.Client{},
	}, nil
}

// Start the message processor
func (m *MessageProcessor) Start() {
	go func() {
		msg := <-m.dataChannel

		subscriptions := m.subscriptionStore.GetSubscriptionsForTopic(msg.Topic)

		for _, s := range subscriptions {
			go func(sub *Subscription) {
				reader := bytes.NewReader(msg.Data)
				req, err := http.NewRequest(http.MethodPost, sub.DestinationURL, reader)

				if err != nil {
					m.logger.Print(err)
					return
				}

				res, err := m.client.Do(req)

				if err != nil {
					m.logger.Print(err)
					return
				}

				if res.StatusCode < 300 {
					sub.Metrics.Success++
				} else {
					sub.Metrics.Failure++
				}

				sub.Save()
			}(&s)
		}
	}()
}

// AddProvider to the processor
func (m *MessageProcessor) AddProvider(provider MessageProvider) error {
	if m.stopped {
		return errors.New("processor is stopped")
	}

	m.providers = append(m.providers, provider)
	provider.OnDataReceived(m.dataChannel)

	return nil
}

// Stop the message processor
func (m *MessageProcessor) Stop() {
	m.stopped = true
	close(m.dataChannel)
}
