package common

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
