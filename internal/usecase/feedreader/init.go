package uFeedreader

import (
	"context"

	m "github.com/jovanfrandika/gochita/domain"
	rCassandra "github.com/jovanfrandika/gochita/internal/repository/cassandra"
	rHttpcall "github.com/jovanfrandika/gochita/internal/repository/httpcall"
)

type usecase struct {
	dbRepo *rCassandra.Repository
	client *rHttpcall.Repository
}

type Usecase interface {
	AddShowEpisodes(ctx context.Context) (err error)
	AddHeadlines(ctx context.Context) (err error)
	AddMangaPosts(ctx context.Context) (err error)
}

func New(dbRepo *rCassandra.Repository, client *rHttpcall.Repository, timeCfg *m.TimeConfig) Usecase {
	return &usecase{
		dbRepo: dbRepo,
		client: client,
	}
}
