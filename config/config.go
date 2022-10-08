package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	m "github.com/jovanfrandika/livechart-notifier/domain"
)

var (
	App          *m.Config
	TimeLocation *time.Location
)

func InitTest() {
	App = &m.Config{
		ChannelID: "random-channel-id",
		RSSUrl:    "https://www.livechart.me/feeds/episodes",
		Timezone:  "Asia/Jakarta",
		Token:     "random-token",
	}

	TimeLocation, _ = time.LoadLocation(App.Timezone)
}

func Init() error {
	if App != nil {
		return nil
	}

	basePath, err := os.Getwd()
	if err != nil {
		return err
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "files", "config.json"))
	if err != nil {
		return err
	}

	App = &m.Config{}
	err = json.Unmarshal(bts, &App)
	if err != nil {
		return err
	}

	TimeLocation, err = time.LoadLocation(App.Timezone)
	if err != nil {
		return err
	}

	return nil
}
