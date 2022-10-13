package rHttpcall

import (
	"time"

	m "github.com/jovanfrandika/livechart-notifier/domain"
)

type repository struct {
	baseUrl      string
	timeLocation *time.Location
}

type Repository interface {
	GetLatestEpisodes() (showMap map[string]m.FeedShow, err error)
}

func New(baseUrl string, timeLocation *time.Location) Repository {
	return &repository{
		baseUrl:      baseUrl,
		timeLocation: timeLocation,
	}
}
