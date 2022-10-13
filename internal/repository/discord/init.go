package rDiscord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jovanfrandika/livechart-notifier/config"
)

type BotClient struct {
	dg *discordgo.Session
}

func New() *BotClient {
	var err error

	dg, err := discordgo.New("Bot " + config.Cfg.Bot.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil
	}

	botClient := &BotClient{
		dg: dg,
	}

	return botClient
}
