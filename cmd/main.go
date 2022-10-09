package main

import (
	"sync"
	"time"

	"github.com/jovanfrandika/livechart-notifier/config"
	bot "github.com/jovanfrandika/livechart-notifier/internal/bot"
	parser "github.com/jovanfrandika/livechart-notifier/internal/parser"
)

func main() {
	config.Init()

	var mtx sync.Mutex
	go func() {
		for {
			mtx.Lock()
			showMap, err := parser.Parse()
			if err == nil {
				bot.NotifyShows(showMap)
			}
			mtx.Unlock()

			time.Sleep(30 * time.Minute)
		}
	}()

	go func() {
		for {
			mtx.Lock()
			parser.CleanShowEpisodes()
			mtx.Unlock()

			time.Sleep(24 * time.Hour)
		}
	}()

	bot.Init()
}
