package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/gochita/domain"
)

func (r *repository) GetSubscriptionsByReferenceId(ctx context.Context, subscriptionType int, referenceId string, isEnabled bool) (channelSubscriptions []m.DbChannelSubscription, err error) {
	channelSubscriptions = []m.DbChannelSubscription{}
	iter := r.session.Query(queryGetSubscriptionsByReferenceId, subscriptionType, referenceId, isEnabled).Iter()
	if err != nil {
		return []m.DbChannelSubscription{}, err
	}
	var channelSubscription m.DbChannelSubscription
	for iter.Scan(&channelSubscription.Id, &channelSubscription.SubscriptionType, &channelSubscription.ReferenceId, &channelSubscription.ContextId, &channelSubscription.IsEnabled) {
		channelSubscriptions = append(channelSubscriptions, channelSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelSubscription{}, err
	}

	return channelSubscriptions, err
}

func (r *repository) GetSubscriptionsByContextId(ctx context.Context, subscriptionType int, contextId string, isEnabled bool) (channelSubscriptions []m.DbChannelSubscription, err error) {
	channelSubscriptions = []m.DbChannelSubscription{}
	iter := r.session.Query(queryGetSubscriptionsByContextId, subscriptionType, contextId, isEnabled).Iter()
	if err != nil {
		return []m.DbChannelSubscription{}, err
	}
	var channelSubscription m.DbChannelSubscription
	for iter.Scan(&channelSubscription.Id, &channelSubscription.SubscriptionType, &channelSubscription.ReferenceId, &channelSubscription.ContextId, &channelSubscription.IsEnabled) {
		channelSubscriptions = append(channelSubscriptions, channelSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelSubscription{}, err
	}

	return channelSubscriptions, err
}

func (r *repository) GetSubscriptions(ctx context.Context, subscriptionType int, isEnabled bool) (channelSubscriptions []m.DbChannelSubscription, err error) {
	channelSubscriptions = []m.DbChannelSubscription{}
	iter := r.session.Query(queryGetSubscriptions, subscriptionType, isEnabled).Iter()
	if err != nil {
		return []m.DbChannelSubscription{}, err
	}
	var channelSubscription m.DbChannelSubscription
	for iter.Scan(&channelSubscription.Id, &channelSubscription.SubscriptionType, &channelSubscription.ReferenceId, &channelSubscription.ContextId, &channelSubscription.IsEnabled) {
		channelSubscriptions = append(channelSubscriptions, channelSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelSubscription{}, err
	}

	return channelSubscriptions, err
}

func (r *repository) GetSubscription(ctx context.Context, subscriptionType int, referenceId, contextId string) (channelSubscription m.DbChannelSubscription, err error) {
	channelSubscription = m.DbChannelSubscription{}
	err = r.session.Query(queryGetSubscription, subscriptionType, referenceId, contextId).Consistency(gocql.One).Scan(&channelSubscription.Id, &channelSubscription.SubscriptionType, &channelSubscription.ReferenceId, &channelSubscription.ContextId, &channelSubscription.IsEnabled)
	if err != nil {
		return m.DbChannelSubscription{}, err
	}

	return channelSubscription, err
}

func (r *repository) CreateSubscription(ctx context.Context, subscriptionType int, referenceId, contextId string) (channelSubscriptionId string, err error) {
	uuid := gocql.TimeUUID()
	err = r.session.Query(queryCreateSubscription, uuid, subscriptionType, referenceId, contextId).Exec()
	if err != nil {
		return "", err
	}
	channelSubscriptionId = uuid.String()
	return channelSubscriptionId, err
}

func (r *repository) ToggleSubscription(ctx context.Context, isEnabled bool, subscriptionType int, referenceId, contextId string) (err error) {
	err = r.session.Query(queryToggleSubscription, isEnabled, subscriptionType, referenceId, contextId).Exec()
	return err
}
