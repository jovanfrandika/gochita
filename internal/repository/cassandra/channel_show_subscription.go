package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetSubscriptionsByReferenceId(ctx context.Context, referenceId string, isEnabled bool) (channelShowSubscriptions []m.DbChannelShowSubscription, err error) {
	channelShowSubscriptions = []m.DbChannelShowSubscription{}
	iter := r.session.Query(queryGetSubscriptionsByReferenceId, referenceId, isEnabled).Iter()
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

func (r *repository) GetSubscriptionsByShowId(ctx context.Context, showId string, isEnabled bool) (channelShowSubscriptions []m.DbChannelShowSubscription, err error) {
	channelShowSubscriptions = []m.DbChannelShowSubscription{}
	iter := r.session.Query(queryGetSubscriptionsByShowId, showId, isEnabled).Iter()
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

func (r *repository) GetSubscription(ctx context.Context, referenceId, showId string) (channelShowSubscription m.DbChannelShowSubscription, err error) {
	channelShowSubscription = m.DbChannelShowSubscription{}
	err = r.session.Query(queryGetSubscription, referenceId, showId).Consistency(gocql.One).Scan(&channelShowSubscription.ReferenceId, &channelShowSubscription.ShowId, &channelShowSubscription.IsEnabled)
	if err != nil {
		return m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscription, err
}

func (r *repository) CreateSubscription(ctx context.Context, referenceId, showId string) (err error) {
	err = r.session.Query(queryCreateSubscription, referenceId, showId).Exec()
	return err
}

func (r *repository) ToggleSubscription(ctx context.Context, isEnabled bool, referenceId, showId string) (err error) {
	err = r.session.Query(queryToggleSubscription, isEnabled, referenceId, showId).Exec()
	return err
}
