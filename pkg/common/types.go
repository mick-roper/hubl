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

	SubscriptionMetrics struct {
		Success int64
		Failure int64
	}

	// SubscriptionStore dols subscription information
	SubscriptionStore interface {
		GetSubscriptionsForTopic(topic string) []Subscription
		PutSubscription(*Subscription) error
		DeleteSubscription(id string) error
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

	return &Subscription{
		ID:             uuid.New().String(),
		Topic:          topic,
		DestinationURL: destinationURL,
	}, nil
}

func (s *Subscription) Save() error {
	return s.store.PutSubscription(s)
}
