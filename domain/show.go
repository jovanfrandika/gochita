package model

type FeedShow struct {
	Title          string
	Thumbnail      string
	Category       string
	Ref            string
	ShowEpisodeMap map[int]FeedShowEpisode
}

type DbShow struct {
	Id             string
	Title          string
	Thumbnail      string
	Category       string
	Ref            string
	ShowEpisodeMap map[string]DbShowEpisode
}
