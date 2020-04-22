package data

import (
	"errors"

	"github.com/mick-roper/hubl/pkg/common"
)

type (
	// MemorySubscriptionStore that uses an in-memory collection
	MemorySubscriptionStore struct {
		subscriptions []common.Subscription
	}
)

// NewMemorySubscriptionStore creates a new subscription store
func NewMemorySubscriptionStore() (*MemorySubscriptionStore, error) {
	return &MemorySubscriptionStore{
		subscriptions: []common.Subscription{},
	}, nil
}

// GetSubscriptionsForTopic gets all subscriptions for a topic
func (s *MemorySubscriptionStore) GetSubscriptionsForTopic(topic string) []common.Subscription {
	result := []common.Subscription{}

	for i := range s.subscriptions {
		if s.subscriptions[i].Topic == topic {
			result = append(result, s.subscriptions[i])
		}
	}

	return result
}

// PutSubscription creates a new subscription
func (s *MemorySubscriptionStore) PutSubscription(subscription *common.Subscription) error {
	if subscription == nil {
		return errors.New("subscription is nil")
	}

	s.DeleteSubscription(subscription)

	s.subscriptions = append(s.subscriptions, *subscription)

	return nil
}

// DeleteSubscription deletes a subscription
func (s *MemorySubscriptionStore) DeleteSubscription(subscription *common.Subscription) error {
	if i := s.findIndex(subscription); i > -1 {
		s.subscriptions = append(s.subscriptions[:i], s.subscriptions[i+1:]...)
	}

	return nil
}

func (s *MemorySubscriptionStore) findIndex(subscription *common.Subscription) (i int) {
	i = -1
	for x, sub := range s.subscriptions {
		if subscription == &sub {
			i = x
			break
		}
	}

	return
}

func (s *MemorySubscriptionStore) Close() {
	// do nothing
}
