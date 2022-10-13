package rCassandra

import (
	"context"
	"log"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

type repository struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

type Repository interface {
	CloseConnection()
	GetShowByTitle(ctx context.Context, title string) (show m.DbShow, err error)
	CreateShow(ctx context.Context, show m.FeedShow) (showId string, err error)
	GetShowEpisodesByShowId(ctx context.Context, showId string) (showEpisodes []m.DbShowEpisode, err error)
	CreateShowEpisode(ctx context.Context, showId string, showEpisode m.FeedShowEpisode) (showEpisodeId string, err error)
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
