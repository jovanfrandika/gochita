package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetHeadlineSubscriptions(ctx context.Context) (channelHeadlineSubscriptions []m.DbChannelHeadlineSubscription, err error) {
	channelHeadlineSubscriptions = []m.DbChannelHeadlineSubscription{}
	iter := r.session.Query(queryGetHeadlineSubscriptions).Iter()
	if err != nil {
		return []m.DbChannelHeadlineSubscription{}, err
	}
	var channelHeadlineSubscription m.DbChannelHeadlineSubscription
	for iter.Scan(&channelHeadlineSubscription.ReferenceId, &channelHeadlineSubscription.IsEnabled) {
		channelHeadlineSubscriptions = append(channelHeadlineSubscriptions, channelHeadlineSubscription)
	}
	if err = iter.Close(); err != nil {
		return []m.DbChannelHeadlineSubscription{}, err
	}

	return channelHeadlineSubscriptions, err
}

func (r *repository) GetHeadlineSubscription(ctx context.Context, referenceId string) (channelHeadlineSubscription m.DbChannelHeadlineSubscription, err error) {
	channelHeadlineSubscription = m.DbChannelHeadlineSubscription{}
	err = r.session.Query(queryGetHeadlineSubscription, referenceId).Consistency(gocql.One).Scan(&channelHeadlineSubscription.ReferenceId, &channelHeadlineSubscription.IsEnabled)
	if err != nil {
		return m.DbChannelHeadlineSubscription{}, err
	}

	return channelHeadlineSubscription, err
}

func (r *repository) ToggleHeadlineSubscription(ctx context.Context, isEnabled bool, referenceId string) (err error) {
	err = r.session.Query(queryToggleHeadlineSubscription, isEnabled, referenceId).Exec()
	return err
}
