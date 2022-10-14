package dBot

import "github.com/bwmarrin/discordgo"

var (
	QUERY = "query"

	COMMAND_SHOW_LIST        = "show-list"
	COMMAND_SHOW_SUBSCRIBE   = "show-subscribe"
	COMMAND_SHOW_UNSUBSCRIBE = "show-unsubscribe"

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        COMMAND_SHOW_LIST,
			Description: "List subscribed shows of this channel",
		},
		{
			Name:        COMMAND_SHOW_SUBSCRIBE,
			Description: "Subscribe show to this channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        QUERY,
					Description: "Show title",
					Required:    true,
				},
			},
		},
		{
			Name:        COMMAND_SHOW_UNSUBSCRIBE,
			Description: "Unsubscribe show to this channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        QUERY,
					Description: "Show title",
					Required:    true,
				},
			},
		},
	}
)
