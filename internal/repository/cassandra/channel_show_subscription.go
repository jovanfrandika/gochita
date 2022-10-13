package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetSubscriptionsByReferenceId(ctx context.Context, referenceId string) (channelShowSubscriptions []m.DbChannelShowSubscription, err error) {
	channelShowSubscriptions = []m.DbChannelShowSubscription{}
	iter := r.session.Query(queryGetSubscriptionsByReferenceId, referenceId).Iter()
	if err != nil {
		return []m.DbChannelShowSubscription{}, err
	}
	var channelShowSubscription m.DbChannelShowSubscription
	for iter.Scan(&channelShowSubscription.ShowId, &channelShowSubscription.ReferenceId, &channelShowSubscription.IsEnabled) {
		channelShowSubscriptions = append(channelShowSubscriptions, channelShowSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscriptions, err
}

func (r *repository) GetSubscriptionsByShowId(ctx context.Context, showId string) (channelShowSubscriptions []m.DbChannelShowSubscription, err error) {
	channelShowSubscriptions = []m.DbChannelShowSubscription{}
	iter := r.session.Query(queryGetSubscriptionsByShowId, showId).Iter()
	if err != nil {
		return []m.DbChannelShowSubscription{}, err
	}
	var channelShowSubscription m.DbChannelShowSubscription
	for iter.Scan(&channelShowSubscription.ShowId, &channelShowSubscription.ReferenceId, &channelShowSubscription.IsEnabled) {
		channelShowSubscriptions = append(channelShowSubscriptions, channelShowSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscriptions, err
}

func (r *repository) GetSubscription(ctx context.Context, showId, referenceId string) (channelShowSubscription m.DbChannelShowSubscription, err error) {
	channelShowSubscription = m.DbChannelShowSubscription{}
	err = r.session.Query(queryGetSubscriptionsByReferenceId, showId, referenceId).Consistency(gocql.One).Scan(&channelShowSubscription.ShowId, &channelShowSubscription.ReferenceId, &channelShowSubscription.IsEnabled)
	if err != nil {
		return m.DbChannelShowSubscription{}, err
	}

	return channelShowSubscription, err
}

func (r *repository) CreateSubscription(ctx context.Context, showId, referenceId string) (err error) {
	err = r.session.Query(queryCreateSubscription, showId, referenceId).Exec()
	return err
}

func (r *repository) ToggleSubscription(ctx context.Context, isEnabled bool, showId, referenceId string) (err error) {
	err = r.session.Query(queryToggleSubscription, isEnabled, showId, referenceId).Exec()
	return err
}
