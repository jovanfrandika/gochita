package model

import "time"

type FeedHeadline struct {
	Title     string
	Thumbnail string
	PubDate   time.Time
	Ref       string
}

type DbHeadline struct {
	Id          string
	Title       string
	Thumbnail   string
	PublishedAt time.Time
	Ref         string
}
