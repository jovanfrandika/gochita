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

	GetShowSubscriptionsByReferenceId(ctx context.Context, referenceId string, isEnabled bool) (channelShowSubscriptions []m.DbChannelShowSubscription, err error)
	GetShowSubscriptionsByShowId(ctx context.Context, showId string, isEnabled bool) (channelShowSubscriptions []m.DbChannelShowSubscription, err error)
	GetShowSubscription(ctx context.Context, referenceId, showId string) (channelShowSubscription m.DbChannelShowSubscription, err error)
	ToggleShowSubscription(ctx context.Context, isEnabled bool, referenceId, showId string) (err error)

	GetShowNotification(ctx context.Context, showEpisodeId, referenceId string) (channelShowEpisodeNotification m.DbChannelShowEpisodeNotification, err error)
	CreateShowNotification(ctx context.Context, showId, referenceId string) (err error)

	GetHeadlinesByRange(ctx context.Context, start, end time.Time) (headlines []m.DbHeadline, err error)
	GetHeadlineByTitle(ctx context.Context, title string) (headline m.DbHeadline, err error)
	CreateHeadline(ctx context.Context, headline m.FeedHeadline) (headlineId string, err error)

	GetHeadlineSubscriptions(ctx context.Context) (channelHeadlineSubscriptions []m.DbChannelHeadlineSubscription, err error)
	GetHeadlineSubscription(ctx context.Context, referenceId string) (channelHeadlineSubscription m.DbChannelHeadlineSubscription, err error)
	ToggleHeadlineSubscription(ctx context.Context, isEnabled bool, referenceId string) (err error)

	GetHeadlineNotification(ctx context.Context, referenceId, headlineId string) (channelHeadlineNotification m.DbChannelHeadlineNotification, err error)
	CreateHeadlineNotification(ctx context.Context, referenceId, headlineId string) (err error)
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
