package rCassandra

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetNotification(ctx context.Context, referenceId, showEpisodeId string) (channelShowEpisodeNotification m.DbChannelShowEpisodeNotification, err error) {
	channelShowEpisodeNotification = m.DbChannelShowEpisodeNotification{}
	err = r.session.Query(queryGetNotification, referenceId, showEpisodeId).Consistency(gocql.One).Scan(&channelShowEpisodeNotification.ReferenceId, &channelShowEpisodeNotification.ShowEpisodeId, &channelShowEpisodeNotification.NotifiedAt)
	return channelShowEpisodeNotification, err
}

func (r *repository) CreateNotification(ctx context.Context, referenceId, showId string) (err error) {
	err = r.session.Query(queryCreateNotification, referenceId, showId, time.Now()).Exec()
	return err
}
