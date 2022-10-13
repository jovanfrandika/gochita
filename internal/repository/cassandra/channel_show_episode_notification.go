package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetNotification(ctx context.Context, showEpisodeId, referenceId string) (channelShowEpisodeNotification m.DbChannelShowEpisodeNotification, err error) {
	channelShowEpisodeNotification = m.DbChannelShowEpisodeNotification{}
	err = r.session.Query(queryCreateNotification, showEpisodeId, referenceId).Consistency(gocql.One).Scan(&channelShowEpisodeNotification.ShowEpisodeId, &channelShowEpisodeNotification.ReferenceId, &channelShowEpisodeNotification.NotifiedAt)
	return channelShowEpisodeNotification, err
}

func (r *repository) CreateNotification(ctx context.Context, showId, referenceId string) (err error) {
	err = r.session.Query(queryCreateSubscription, showId, referenceId).Exec()
	return err
}
