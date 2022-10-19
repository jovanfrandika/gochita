package uLivechart

import (
	"context"

	rCassandra "github.com/jovanfrandika/livechart-notifier/internal/repository/cassandra"
	rHttpcall "github.com/jovanfrandika/livechart-notifier/internal/repository/httpcall"
)

type usecase struct {
	dbRepo *rCassandra.Repository
	client *rHttpcall.Repository
}

type Usecase interface {
	AddShowEpisodes(ctx context.Context) (err error)
	AddHeadlines(ctx context.Context) (err error)
}

func New(dbRepo *rCassandra.Repository, client *rHttpcall.Repository) Usecase {
	return &usecase{
		dbRepo: dbRepo,
		client: client,
	}
}
