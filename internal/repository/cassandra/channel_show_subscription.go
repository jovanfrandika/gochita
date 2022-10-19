package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetShowSubscriptionsByReferenceId(ctx context.Context, referenceId string, isEnabled bool) (channelShowSubscriptions []m.DbChannelShowSubscription, err error) {
	channelShowSubscriptions = []m.DbChannelShowSubscription{}
	iter := r.session.Query(queryGetShowSubscriptionsByReferenceId, referenceId, isEnabled).Iter()
	if err != nil {
		return []m.DbChannelShowSubscription{}, err
	}
	var channelShowSubscription m.DbChannelShowSubscription
	for iter.Scan(&channelShowSubscription.ReferenceId, &channelShowSubscription.ShowId, &channelShowSubscription.IsEnabled) {
		channelShowSubscriptions = append(channelShowSubscriptions, channelShowSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscriptions, err
}

func (r *repository) GetShowSubscriptionsByShowId(ctx context.Context, showId string, isEnabled bool) (channelShowSubscriptions []m.DbChannelShowSubscription, err error) {
	channelShowSubscriptions = []m.DbChannelShowSubscription{}
	iter := r.session.Query(queryGetShowSubscriptionsByShowId, showId, isEnabled).Iter()
	if err != nil {
		return []m.DbChannelShowSubscription{}, err
	}
	var channelShowSubscription m.DbChannelShowSubscription
	for iter.Scan(&channelShowSubscription.ReferenceId, &channelShowSubscription.ShowId, &channelShowSubscription.IsEnabled) {
		channelShowSubscriptions = append(channelShowSubscriptions, channelShowSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscriptions, err
}

func (r *repository) GetShowSubscription(ctx context.Context, referenceId, showId string) (channelShowSubscription m.DbChannelShowSubscription, err error) {
	channelShowSubscription = m.DbChannelShowSubscription{}
	err = r.session.Query(queryGetShowSubscription, referenceId, showId).Consistency(gocql.One).Scan(&channelShowSubscription.ReferenceId, &channelShowSubscription.ShowId, &channelShowSubscription.IsEnabled)
	if err != nil {
		return m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscription, err
}

func (r *repository) ToggleShowSubscription(ctx context.Context, isEnabled bool, referenceId, showId string) (err error) {
	err = r.session.Query(queryToggleShowSubscription, isEnabled, referenceId, showId).Exec()
	return err
}
