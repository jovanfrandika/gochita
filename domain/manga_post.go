package model

import "time"

type FeedMangaPost struct {
	Title   string
	PubDate time.Time
	Ref     string
}

type DbMangaPost struct {
	Id          string
	Title       string
	PublishedAt time.Time
	Ref         string
}
