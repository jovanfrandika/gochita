package model

import "time"

type FeedShowEpisode struct {
	Num     int
	Ref     string
	PubDate time.Time
}

type FeedShow struct {
	Title          string
	Thumbnail      string
	Category       string
	Ref            string
	ShowEpisodeMap map[int]FeedShowEpisode
}

type DbShowEpisode struct {
	Id      string
	ShowId  string
	Num     int
	PubDate time.Time
}

type DbShow struct {
	Id             string
	Title          string
	Thumbnail      string
	Category       string
	Ref            string
	ShowEpisodeMap map[string]DbShowEpisode
}
