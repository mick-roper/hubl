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
