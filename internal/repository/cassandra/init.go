package rCassandra

import (
	"context"
	"log"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/gochita/domain"
)

type repository struct {
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

	GetSubscriptionsByReferenceId(ctx context.Context, subscriptionType int, referenceId string, isEnabled bool) (channelSubscriptions []m.DbChannelSubscription, err error)
	GetSubscriptionsByContextId(ctx context.Context, subscriptionType int, contextId string, isEnabled bool) (channelSubscriptions []m.DbChannelSubscription, err error)
	GetSubscriptions(ctx context.Context, subscriptionType int, isEnabled bool) (channelSubscriptions []m.DbChannelSubscription, err error)
	GetSubscription(ctx context.Context, subscriptionType int, referenceId, contextId string) (channelSubscription m.DbChannelSubscription, err error)
	CreateSubscription(ctx context.Context, subscriptionType int, referenceId, contextId string) (channelSubscriptionId string, err error)
	ToggleSubscriptions(ctx context.Context, isEnabled bool, subscriptionType int, referenceId string, contextIds []gocql.UUID) (err error)
	ToggleSubscription(ctx context.Context, isEnabled bool, subscriptionType int, referenceId, contextId string) (err error)

	GetNotification(ctx context.Context, channelSubscriptionId string) (channelNotification m.DbChannelNotification, err error)
	CreateNotification(ctx context.Context, channelSubscriptionId string) (err error)

	GetHeadlinesByRange(ctx context.Context, start, end time.Time) (headlines []m.DbHeadline, err error)
	GetHeadlineByTitle(ctx context.Context, title string) (headline m.DbHeadline, err error)
	CreateHeadline(ctx context.Context, headline m.FeedHeadline) (headlineId string, err error)

	GetMangaPostsByRange(ctx context.Context, start, end time.Time) (mangaPosts []m.DbMangaPost, err error)
	GetMangaPostByTitle(ctx context.Context, title string) (mangaPost m.DbMangaPost, err error)
	CreateMangaPost(ctx context.Context, mangaPost m.FeedMangaPost) (mangaPostId string, err error)
}

func New(clusters []string, keyspaceName string, timeout int) Repository {
	cluster := gocql.NewCluster(clusters...)
	cluster.Keyspace = keyspaceName
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Duration(timeout) * time.Second

	r := &repository{}
	var err error
	r.session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	return r
}
