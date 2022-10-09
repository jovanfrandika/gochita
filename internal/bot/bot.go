package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jovanfrandika/livechart-notifier/config"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

var (
	dg *discordgo.Session
)

func Init() {
	var err error
	dg, err = discordgo.New("Bot " + config.App.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer dg.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func NotifyShows(showMap map[string]*m.Show) {
	if dg != nil {
		for showName, episode := range showMap {
			for episodeNum, episodeDetail := range episode.ShowEpisodeMap {
				if !episodeDetail.IsNotified {
					content := fmt.Sprintf("***Reminder!!!***\nTitle: %v\nEpisode: %v\nPublished Date: %v\n", showName, episodeNum, episodeDetail.PubDate.Format(time.RFC850))
					dg.ChannelMessageSendComplex(config.App.ChannelID, &discordgo.MessageSend{
						Content: content,
						Embeds: []*discordgo.MessageEmbed{
							{
								Image: &discordgo.MessageEmbedImage{
									URL: episode.Enclosure.Url,
								},
								Type: discordgo.EmbedTypeImage,
							},
						},
					})
					showMap[showName].ShowEpisodeMap[episodeNum].IsNotified = true
				}
			}
		}
	}
}
