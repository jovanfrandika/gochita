package rCassandra

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetHeadlineNotification(ctx context.Context, referenceId, headlineId string) (channelHeadlineNotification m.DbChannelHeadlineNotification, err error) {
	channelHeadlineNotification = m.DbChannelHeadlineNotification{}
	err = r.session.Query(queryGetHeadlineNotification, referenceId, headlineId).Consistency(gocql.One).Scan(&channelHeadlineNotification.ReferenceId, &channelHeadlineNotification.HeadlineId, &channelHeadlineNotification.NotifiedAt)
	return channelHeadlineNotification, err
}

func (r *repository) CreateHeadlineNotification(ctx context.Context, referenceId, headlineId string) (err error) {
	err = r.session.Query(queryCreateHeadlineNotification, referenceId, headlineId, time.Now()).Exec()
	return err
}
