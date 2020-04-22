package common

import (
	"errors"

	"github.com/google/uuid"
)

type (
	// Subscription that describes a topic and a destination
	Subscription struct {
		ID             string
		Topic          string
		DestinationURL string
		Metrics        SubscriptionMetrics
		store          SubscriptionStore
	}

	// SubscriptionMetrics that describe what has previously happened with this subscription
	SubscriptionMetrics struct {
		Success int64
		Failure int64
	}

	// SubscriptionStore dols subscription information
	SubscriptionStore interface {
		GetAll() []Subscription
		GetSubscriptionsForTopic(topic string) []Subscription
		PutSubscription(*Subscription) error
		DeleteSubscription(*Subscription) error
		Close()
	}
)

// NewSubscription creates a new Subscription object
func NewSubscription(topic, destinationURL string, store SubscriptionStore) (*Subscription, error) {
	if topic == "" {
		return nil, errors.New("topic cannot be empty")
	}

	if destinationURL == "" {
		return nil, errors.New("desinationURL cannot be empty")
	}

	if store == nil {
		return nil, errors.New("store cannot be nil")
	}

	return &Subscription{
		ID:             uuid.New().String(),
		Topic:          topic,
		DestinationURL: destinationURL,
		store:          store,
	}, nil
}

// Save the subscription
func (s *Subscription) Save() error {
	return s.store.PutSubscription(s)
}
