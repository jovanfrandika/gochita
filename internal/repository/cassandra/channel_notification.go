package rCassandra

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/gochita/domain"
)

func (r *repository) GetNotification(ctx context.Context, channelSubscriptionId string) (channelNotification m.DbChannelNotification, err error) {
	channelNotification = m.DbChannelNotification{}
	err = r.session.Query(queryGetNotification, channelSubscriptionId).Consistency(gocql.One).Scan(&channelNotification.ChannelSubscriptionId, &channelNotification.NotifiedAt)
	return channelNotification, err
}

func (r *repository) CreateNotification(ctx context.Context, channelSubscriptionId string) (err error) {
	err = r.session.Query(queryCreateNotification, channelSubscriptionId, time.Now()).Exec()
	return err
}
