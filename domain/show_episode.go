package model

import "time"

type FeedShowEpisode struct {
	Num     int
	Ref     string
	PubDate time.Time
}

type DbShowEpisode struct {
	Id      string
	ShowId  string
	Num     int
	PubDate time.Time
}
