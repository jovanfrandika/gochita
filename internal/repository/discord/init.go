package rDiscord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type discordBotRepo struct {
	dg *discordgo.Session
}

type DiscordBotRepo interface {
	AddHandler(handler interface{})
	Connect() (err error)
	Close()
	SendMsgToChannel(channelId string, msg *discordgo.MessageSend)
	RegisterCommand(cmd *discordgo.ApplicationCommand) (ccmd *discordgo.ApplicationCommand, err error)
	UnregisterCommand(cmd *discordgo.ApplicationCommand) (err error)
}

func New(token string) DiscordBotRepo {
	var err error

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	return &discordBotRepo{
		dg: dg,
	}
}
