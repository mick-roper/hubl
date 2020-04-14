package common

type Subscription struct {
	Topic string
}

type SubscriptionService interface {
	Add(*Subscription) error
	GetForTopic(topic string) ([]Subscription, error)
}
