package memory

import (
	"errors"

	"github.com/mick-roper/hubl/pkg/common"
)

type (
	// SubscriptionStore that uses an in-memory collection
	SubscriptionStore struct {
		subscriptions []common.Subscription
		itemCount     int
		increment     int
	}
)

// NewSubscriptionStore creates a new subscription store
func NewSubscriptionStore(capacity int) (*SubscriptionStore, error) {
	if capacity < 1 {
		return nil, errors.New("capacity must be at least 1")
	}

	return &SubscriptionStore{
		subscriptions: make([]Subscription, capacity),
		itemCount:     0,
		increment:     capacity,
	}, nil
}

// GetSubscriptionsForTopic gets all subscriptions for a topic
func (s *SubscriptionStore) GetSubscriptionsForTopic(topic string) []Subscription {
	result := []Subscription{}

	for i := range s.subscriptions {
		if s.subscriptions[i].Topic == topic {
			result = append(result, s.subscriptions[i])
		}
	}

	return result
}

// AddSubscription creates a new subscription
func (s *SubscriptionStore) AddSubscription(s common.Subscription) error {
	return errors.New("not implemented")
}

// DeleteSubscription deletes a subscription
func (s *SubscriptionStore) DeleteSubscription(id string) error {
	return errors.New("not implemented")
}
