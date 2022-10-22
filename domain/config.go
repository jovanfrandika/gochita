package model

import "time"

type BotConfig struct {
	Token string `json:"token"`
}

type DBConfig struct {
	KeyspaceName string   `json:"keyspaceName"`
	Clusters     []string `json:"clusters"`
}

type LiveChartConfig struct {
	BaseUrl            string `json:"baseUrl"`
	UriLatestEpisodes  string `json:"uriLatestEpisodes"`
	UriLatestHeadlines string `json:"uriLatestHeadlines"`
}

type RedditConfig struct {
	BaseUrl             string `json:"baseUrl"`
	UriLatestMangaPosts string `json:"uriLatestMangaPosts"`
}

type TimeConfig struct {
	Timezone                string `json:"timezone"`
	TimeLocation            *time.Location
	DefaultTimeout          int `json:"defaultTimeout"`
	NotifyTimeout           int `json:"notifyTimeout"`
	SetCommandsTimeout      int `json:"setCommandsTimeout"`
	NotifyShowsInterval     int `json:"notifyShowsInterval"`
	NotifyHeadlinesInterval int `json:"notfiyHeadlinesInterval"`
	NotifyMangasInterval    int `json:"notifyMangasInterval"`
	AddShowsInterval        int `json:"addShowsInterval"`
	AddHeadlinesInterval    int `json:"addHeadlinesInterval"`
	AddMangasInterval       int `json:"addMangasInterval"`
}

type Config struct {
	Bot       BotConfig       `json:"bot"`
	DB        DBConfig        `json:"db"`
	LiveChart LiveChartConfig `json:"liveChart"`
	Reddit    RedditConfig    `json:"reddit"`
	Time      TimeConfig      `json:"time"`
}
