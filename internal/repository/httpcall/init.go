package rHttpcall

import (
	"time"

	m "github.com/jovanfrandika/gochita/domain"
)

type repository struct {
	baseUrl      string
	timeLocation *time.Location
}

type Repository interface {
	GetLatestEpisodes() (showMap map[string]m.FeedShow, err error)
	GetLatestHeadlines() (headlineMap map[string]m.FeedHeadline, err error)
}

func New(baseUrl string, timeLocation *time.Location) Repository {
	return &repository{
		baseUrl:      baseUrl,
		timeLocation: timeLocation,
	}
}
