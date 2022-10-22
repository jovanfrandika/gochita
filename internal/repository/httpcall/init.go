package rHttpcall

import (
	"crypto/tls"
	"net/http"

	m "github.com/jovanfrandika/gochita/domain"
)

type repository struct {
	livechartCfg *m.LiveChartConfig
	redditCfg    *m.RedditConfig
	timeCfg      *m.TimeConfig
	httpClient   http.Client
}

type Repository interface {
	GetLatestEpisodes() (showMap map[string]m.FeedShow, err error)
	GetLatestHeadlines() (headlineMap map[string]m.FeedHeadline, err error)
	GetLatestMangaPosts() (mangaPostMap map[string]m.FeedMangaPost, err error)
}

func New(livechartCfg *m.LiveChartConfig, redditCfg *m.RedditConfig, timeCfg *m.TimeConfig) Repository {
	return &repository{
		livechartCfg: livechartCfg,
		redditCfg:    redditCfg,
		timeCfg:      timeCfg,
		httpClient: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{},
			},
		},
	}
}
