package memory

import (
	"errors"

	"github.com/mick-roper/hubl/pkg/common"
)

type (
	// SubscriptionStore that uses an in-memory collection
	SubscriptionStore struct {
		subscriptions []common.Subscription
	}
)

// NewSubscriptionStore creates a new subscription store
func NewSubscriptionStore() (*SubscriptionStore, error) {
	return &SubscriptionStore{
		subscriptions: []common.Subscription{},
	}, nil
}

// GetSubscriptionsForTopic gets all subscriptions for a topic
func (s *SubscriptionStore) GetSubscriptionsForTopic(topic string) []Subscription {
	result := []common.Subscription{}

	for i := range s.subscriptions {
		if s.subscriptions[i].Topic == topic {
			result = append(result, s.subscriptions[i])
		}
	}

	return result
}

// AddSubscription creates a new subscription
func (s *SubscriptionStore) AddSubscription(s *common.Subscription) error {
	if s == nil {
		return errors.New("subscription is nil")
	}

	s.subscriptions = append(s.subscriptions, &s)
}

// DeleteSubscription deletes a subscription
func (s *SubscriptionStore) DeleteSubscription(id string) error {
	for i := range s.subscriptions {
		if s.subscriptions[i].ID == id {
			s.subscriptions = append(s.subscriptions[:i], s.subscriptions[i+1:]...)
			return
		}
	}
}
