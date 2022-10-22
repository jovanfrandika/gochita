package rHttpcall

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	m "github.com/jovanfrandika/gochita/domain"
)

func (r *repository) GetLatestEpisodes() (showMap map[string]m.FeedShow, err error) {
	resp, err := r.httpClient.Get(r.livechartCfg.BaseUrl + r.livechartCfg.UriLatestEpisodes)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf(LABEL_STATUS_CODE, resp.StatusCode))
	}

	rss := m.RssFeed{}
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
		var showEpisodeNum int
		showEpisodeArr := regex.FindAllString(item.Title, 1)
		if len(showEpisodeArr) > 0 {
			showEpisodeStr := strings.Trim(showEpisodeArr[0], "#")
			showEpisodeNum, err = strconv.Atoi(showEpisodeStr)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
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
				PubDate: pubDate.In(r.timeCfg.TimeLocation),
			}
		}
	}

	return showMap, nil
}

func (r *repository) GetLatestHeadlines() (headlineMap map[string]m.FeedHeadline, err error) {
	resp, err := r.httpClient.Get(r.livechartCfg.BaseUrl + r.livechartCfg.UriLatestHeadlines)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf(LABEL_STATUS_CODE, resp.StatusCode))
	}

	rss := m.RssFeed{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return nil, err
	}

	headlineMap = map[string]m.FeedHeadline{}
	for _, item := range rss.Channel.Items {
		if _, exists := headlineMap[item.Guid]; !exists {
			pubDate, _ := time.Parse(time.RFC1123Z, item.PubDate)

			headlineMap[item.Guid] = m.FeedHeadline{
				Title:     item.Title,
				Thumbnail: item.Enclosure.Url,
				Ref:       item.Link,
				PubDate:   pubDate.In(r.timeCfg.TimeLocation),
			}
		}
	}

	return headlineMap, nil
}
