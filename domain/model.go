package model

import "time"

type Config struct {
	ChannelID string `json:channelId`
	RSSUrl    string `json:rssUrl`
	Timezone  string `json:timezone`
	Token     string `json:token`
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	Guid      string    `xml:"guid"`
	Enclosure Enclosure `xml:"enclosure"`
	PubDate   string    `xml:"pubDate"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type ShowEpisode struct {
	Episode    string
	Link       string
	IsNotified bool
	PubDate    time.Time
}

type Show struct {
	Title          string
	Category       string
	Link           string
	Enclosure      Enclosure
	ShowEpisodeMap map[string]*ShowEpisode
}
