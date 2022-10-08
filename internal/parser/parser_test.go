package parser

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jovanfrandika/livechart-notifier/config"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func TestParse(t *testing.T) {
	config.InitTest()

	var mockServer *httptest.Server

	tests := []struct {
		name     string
		want     map[string]*m.Show
		wantErr  bool
		mockFunc func()
	}{
		{
			name: "positive case",
			want: map[string]*m.Show{
				"BLUELOCK Additional Time! ": {
					Title:    "BLUELOCK Additional Time! ",
					Category: "Anime/OVAs/Episodes",
					Link:     "https://www.livechart.me/anime/11548",
					Enclosure: m.Enclosure{
						Url:    "https://u.livechart.me/anime/11548/poster_image/741a0e469f0d72b6b55d111c12f0fbc6.jpg?style=small&amp;format=jpg",
						Length: 0,
						Type:   "image/jpeg",
					},
					ShowEpisodeMap: map[string]*m.ShowEpisode{
						"#1": {
							Episode:    "#1",
							Link:       "https://www.livechart.me/anime/11548#episode139723:1665248400",
							IsNotified: false,
							PubDate: func() time.Time {
								date, _ := time.Parse(time.RFC1123Z, "Sat, 08 Oct 2022 17:00:00 +0000")
								return date.In(config.TimeLocation)
							}(),
						},
					},
				},
			},
			wantErr: false,
			mockFunc: func() {
				mockServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					res.Write([]byte(`
							<?xml version="1.0" encoding="utf-8" ?>
							<rss version="2.0"
									xmlns:atom="http://www.w3.org/2005/Atom"
									xmlns:media="http://search.yahoo.com/mrss/"
									xmlns:content="http://purl.org/rss/1.0/modules/content/">
								<channel>  <title>LiveChart.me - Recent Anime Episodes</title>
								<link>https://www.livechart.me/schedule/all</link>
								<lastBuildDate>Sat, 08 Oct 2022 17:00:00 +0000</lastBuildDate>
								<webMaster>mike@livechart.me (Michael Millard)</webMaster>
								<atom:link href="https://www.livechart.me/feeds/episodes" rel="self" type="application/rss+xml" />
								<description>Feed of anime episodes that have aired in the last 24 hours</description>
									<item>
										<guid>https://www.livechart.me/anime/11548#episode139723:1665248400</guid>
										<link>https://www.livechart.me/anime/11548</link>
										<title>BLUELOCK Additional Time! #1</title>
										<pubDate>Sat, 08 Oct 2022 17:00:00 +0000</pubDate>
										<category>Anime/OVAs/Episodes</category>

										<enclosure url="https://u.livechart.me/anime/11548/poster_image/741a0e469f0d72b6b55d111c12f0fbc6.jpg?style=small&amp;format=jpg" length="0" type="image/jpeg" />
										<media:thumbnail url="https://u.livechart.me/anime/11548/poster_image/741a0e469f0d72b6b55d111c12f0fbc6.jpg?style=small&amp;format=jpg" width="175" height="250" />
									</item>
							</channel>
							</rss>
					`))
				}))
				config.App.RSSUrl = mockServer.URL + "/"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
