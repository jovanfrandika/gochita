package model

type DbChannelSubscription struct {
	Id               string
	ReferenceId      string
	ContextId        string
	SubscriptionType int
	IsEnabled        bool
}
