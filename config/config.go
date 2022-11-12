package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	m "github.com/jovanfrandika/gochita/domain"
)

func getenvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal(fmt.Sprintf("%s is empty", key))
	}
	return v
}

func getenvInt(key string) int {
	s := getenvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return v
}

func Init() (cfg *m.Config) {
	cfg = &m.Config{}

	cfg.Bot.Token = getenvStr("TOKEN")

	cfg.DB.KeyspaceName = getenvStr("KEYSPACE_NAME")
	cfg.DB.Clusters = []string{getenvStr("CLUSTER")}
	cfg.DB.Timeout = getenvInt("TIMEOUT")

	cfg.LiveChart.BaseUrl = getenvStr("LIVECHART_BASE_URL")
	cfg.LiveChart.UriLatestEpisodes = getenvStr("LIVECHART_LATEST_EPISODES_URI")
	cfg.LiveChart.UriLatestHeadlines = getenvStr("LIVECHART_LATEST_HEADLINES_URI")

	cfg.Reddit.BaseUrl = getenvStr("REDDIT_BASE_URL")
	cfg.Reddit.UriLatestMangaPosts = getenvStr("REDDIT_LATEST_MANGA_POSTS_URI")

	cfg.Time.Timezone = getenvStr("TIMEZONE")
	cfg.Time.DefaultTimeout = getenvInt("DEFAULT_TIMEOUT")
	cfg.Time.NotifyTimeout = getenvInt("NOTIFY_TIMEOUT")
	cfg.Time.SetCommandsTimeout = getenvInt("NOTIFY_TIMEOUT")
	cfg.Time.NotifyShowsInterval = getenvInt("NOTIFY_SHOWS_INTERVAL")
	cfg.Time.NotifyHeadlinesInterval = getenvInt("NOTIFY_HEADLINES_INTERVAL")
	cfg.Time.NotifyMangasInterval = getenvInt("NOTIFY_MANGAS_INTERVAL")
	cfg.Time.AddShowsInterval = getenvInt("ADD_SHOWS_INTERVAL")
	cfg.Time.AddHeadlinesInterval = getenvInt("ADD_HEADLINES_INTERVAL")
	cfg.Time.AddMangasInterval = getenvInt("ADD_MANGAS_INTERVAL")

	var err error
	cfg.Time.TimeLocation, err = time.LoadLocation(cfg.Time.Timezone)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
