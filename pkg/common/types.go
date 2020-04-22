package common

import "errors"

type (
	// Subscription that describes a topic and a destination
	Subscription struct {
		ID             string
		Topic          string
		DestinationURL string
	}

	// SubscriptionStore dols subscription information
	SubscriptionStore interface {
		GetSubscriptionsForTopic(topic string) []Subscription
		AddSubscription(Subscription) error
		DeleteSubscription(id string) error
	}
)

// NewSubscription creates a new Subscription object
func NewSubscription(topic, destinationURL string) (*Subscription, error) {
	if topic == "" {
		return nil, errors.New("topic cannot be empty")
	}

	if destinationURL == "" {
		return nil, errors.New("desinationURL cannot be empty")
	}

	return &Subscription{
		Topic:          topic,
		DestinationURL: destinationURL,
	}, nil
}
