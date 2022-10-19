package dBot

import "github.com/bwmarrin/discordgo"

var (
	QUERY = "query"

	COMMAND_SHOW_LIST                = "show-list"
	COMMAND_SHOW_SUBSCRIBE_ALL       = "show-subscribe-all"
	COMMAND_SHOW_UNSUBSCRIBE_ALL     = "show-unsubscribe-all"
	COMMAND_SHOW_SUBSCRIBE           = "show-subscribe"
	COMMAND_SHOW_UNSUBSCRIBE         = "show-unsubscribe"
	COMMAND_HEADLINE_SUBSCRIBE_ALL   = "headline-subscribe-all"
	COMMAND_HEADLINE_UNSUBSCRIBE_ALL = "headline-unsubscribe-all"

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        COMMAND_SHOW_LIST,
			Description: "List subscribed shows of this channel",
		},
		{
			Name:        COMMAND_SHOW_SUBSCRIBE_ALL,
			Description: "Subscribe to new shows",
		},
		{
			Name:        COMMAND_SHOW_UNSUBSCRIBE_ALL,
			Description: "Unsubscribe to new shows",
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
		{
			Name:        COMMAND_HEADLINE_SUBSCRIBE_ALL,
			Description: "Subscribe to new headlines",
		},
		{
			Name:        COMMAND_HEADLINE_UNSUBSCRIBE_ALL,
			Description: "Unsubscribe to new headlines",
		},
	}
)
