package parser

import (
	"encoding/xml"
	"net/http"
	"regexp"
	"time"

	"github.com/jovanfrandika/livechart-notifier/config"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

const (
	threeDays = 3 * 24 * time.Hour
)

var (
	showMap map[string]*m.Show = map[string]*m.Show{}
)

func Parse() (map[string]*m.Show, error) {
	resp, err := http.Get(config.App.RSSUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rss := m.Rss{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile(`(#\w+)`)
	if err != nil {
		return nil, err
	}

	for _, item := range rss.Channel.Items {
		showEpisodeArr := regex.FindAllString(item.Title, 1)
		if len(showEpisodeArr) < 0 {
			continue
		}
		showEpisode := showEpisodeArr[0]
		showName := regex.ReplaceAllString(item.Title, "")

		if _, exists := showMap[showName]; !exists {
			showMap[showName] = &m.Show{
				Title:          showName,
				Link:           item.Link,
				Enclosure:      item.Enclosure,
				ShowEpisodeMap: map[string]*m.ShowEpisode{},
			}
		}

		if _, exists := showMap[showName].ShowEpisodeMap[showEpisode]; !exists {
			pubDate, _ := time.Parse(time.RFC1123Z, item.PubDate)

			showMap[showName].ShowEpisodeMap[showEpisode] = &m.ShowEpisode{
				Episode:    showEpisode,
				Link:       item.Guid,
				IsNotified: false,
				PubDate:    pubDate.In(config.TimeLocation),
			}
		}
	}

	return showMap, nil
}

func CleanShowEpisodes() {
	now := time.Now().In(config.TimeLocation)

	for _, episode := range showMap {
		for k, v := range episode.ShowEpisodeMap {
			expiresAt := v.PubDate.Add(threeDays)
			if now.After(expiresAt) {
				delete(episode.ShowEpisodeMap, k)
			}
		}
	}
}
