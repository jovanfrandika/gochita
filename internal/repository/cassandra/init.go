package rCassandra

import (
	"context"
	"log"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

type repository struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

type Repository interface {
	CloseConnection()
	GetShowById(ctx context.Context, id string) (show m.DbShow, err error)
	GetShowByTitle(ctx context.Context, title string) (show m.DbShow, err error)
	CreateShow(ctx context.Context, show m.FeedShow) (showId string, err error)

	GetShowEpisodesByShowId(ctx context.Context, showId string) (showEpisodes []m.DbShowEpisode, err error)
	GetShowEpisodesByRange(ctx context.Context, start, end time.Time) (showEpisodes []m.DbShowEpisode, err error)
	CreateShowEpisode(ctx context.Context, showId string, showEpisode m.FeedShowEpisode) (showEpisodeId string, err error)

	GetSubscriptionsByReferenceId(ctx context.Context, referenceId string) (channelShowSubscriptions []m.DbChannelShowSubscription, err error)
	GetSubscriptionsByShowId(ctx context.Context, showId string) (channelShowSubscriptions []m.DbChannelShowSubscription, err error)
	GetSubscription(ctx context.Context, referenceId, showId string) (channelShowSubscription m.DbChannelShowSubscription, err error)
	CreateSubscription(ctx context.Context, referenceId, showId string) (err error)
	ToggleSubscription(ctx context.Context, isEnabled bool, referenceId, showId string) (err error)

	GetNotification(ctx context.Context, showEpisodeId, referenceId string) (channelShowEpisodeNotification m.DbChannelShowEpisodeNotification, err error)
	CreateNotification(ctx context.Context, showId, referenceId string) (err error)
}

func New(clusters []string, keyspaceName string) Repository {
	r := &repository{
		cluster: gocql.NewCluster(clusters...),
	}
	r.cluster.Keyspace = keyspaceName
	r.cluster.Consistency = gocql.Quorum

	var err error
	r.session, err = r.cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	return r
}
