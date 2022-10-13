package rHttpcall

import (
	"encoding/xml"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	m "github.com/jovanfrandika/livechart-notifier/domain"
)

const (
	threeDays = 3 * 24 * time.Hour

	uriLatestEpisodes = "/feeds/episodes"
)

func (r *repository) GetLatestEpisodes() (showMap map[string]m.FeedShow, err error) {
	resp, err := http.Get(r.baseUrl + uriLatestEpisodes)
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

	showMap = map[string]m.FeedShow{}
	for _, item := range rss.Channel.Items {
		showEpisodeArr := regex.FindAllString(item.Title, 1)
		if len(showEpisodeArr) < 0 {
			continue
		}

		showEpisodeStr := strings.Trim(showEpisodeArr[0], "#")
		showEpisodeNum, err := strconv.Atoi(showEpisodeStr)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		showName := strings.TrimSpace(regex.ReplaceAllString(item.Title, ""))

		if _, exists := showMap[showName]; !exists {
			showMap[showName] = m.FeedShow{
				Title:          showName,
				Thumbnail:      item.Enclosure.Url,
				Category:       item.Desc,
				Ref:            item.Link,
				ShowEpisodeMap: map[int]m.FeedShowEpisode{},
			}
		}

		if _, exists := showMap[showName].ShowEpisodeMap[showEpisodeNum]; !exists {
			pubDate, _ := time.Parse(time.RFC1123Z, item.PubDate)

			showMap[showName].ShowEpisodeMap[showEpisodeNum] = m.FeedShowEpisode{
				Num:     showEpisodeNum,
				Ref:     item.Guid,
				PubDate: pubDate.In(r.timeLocation),
			}
		}
	}

	return showMap, nil
}
