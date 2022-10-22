package rHttpcall

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	m "github.com/jovanfrandika/gochita/domain"
)

func (r *repository) GetLatestMangaPosts() (mangaPostMap map[string]m.FeedMangaPost, err error) {
	resp, err := r.httpClient.Get(r.redditCfg.BaseUrl + r.redditCfg.UriLatestMangaPosts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf(LABEL_STATUS_CODE, resp.StatusCode))
	}

	rss := m.AtomFeed{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return nil, err
	}

	discRegex, err := regexp.Compile(`(\[DISC\])`)
	if err != nil {
		return nil, err
	}

	hrefRegex, err := regexp.Compile(`href="(.+?)"`)
	if err != nil {
		return nil, err
	}

	mangaPostMap = map[string]m.FeedMangaPost{}
	for _, entry := range rss.Entries {
		mangaPostTitleArr := discRegex.FindAllString(entry.Title, 1)
		if len(mangaPostTitleArr) < 1 {
			continue
		}

		mangaTitle := strings.TrimSpace(discRegex.ReplaceAllString(entry.Title, ""))

		if _, exists := mangaPostMap[mangaTitle]; !exists {
			hrefArr := hrefRegex.FindAllStringSubmatch(entry.Content.Content, -1)
			if len(hrefArr) < 0 {
				continue
			}

			var ref string
			for _, match := range hrefArr {
				href := match[1]
				if !strings.Contains(href, r.redditBaseUrl) {
					ref = href
					break
				}
			}
			if ref == "" {
				continue
			}
			pubDate, err := time.Parse("2006-01-02T15:04:05Z07:00", entry.Published)
			if err != nil {
				return nil, err
			}

			mangaPostMap[mangaTitle] = m.FeedMangaPost{
				Title:   mangaTitle,
				Ref:     ref,
				PubDate: pubDate.In(r.timeCfg.TimeLocation),
			}
		}
	}

	return mangaPostMap, nil
}
