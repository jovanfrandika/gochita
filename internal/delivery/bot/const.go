package dBot

import "github.com/bwmarrin/discordgo"

var (
	QUERY = "query"

	SUBCOMMAND_LIST        = "list"
	SUBCOMMAND_ALL         = "ALL"
	SUBCOMMAND_NEW         = "new"
	SUBCOMMAND_ONE         = "one"
	SUBCOMMAND_SUBSCRIBE   = "subscribe"
	SUBCOMMAND_UNSUBSCRIBE = "unsubscribe"

	COMMAND_SHOW       = "show"
	COMMAND_HEADLINE   = "headline"
	COMMAND_MANGA_POST = "manga"

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        COMMAND_SHOW,
			Description: "Show related commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        SUBCOMMAND_LIST,
					Description: "List subscribed shows of this channel",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        SUBCOMMAND_SUBSCRIBE,
					Description: "Subscribe show to this channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_ONE,
							Description: "Subscribe to a show",
							Required:    false,
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
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_NEW,
							Description: "Subscribe to new shows",
							Required:    false,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        SUBCOMMAND_UNSUBSCRIBE,
					Description: "Unsubscribe show to this channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_ONE,
							Description: "Unsubscribe to a show",
							Required:    false,
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
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_NEW,
							Description: "Unsubscribe to new shows",
							Required:    false,
						},
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_ALL,
							Description: "Unsubscribe to all shows",
							Required:    false,
						},
					},
				},
			},
		},
		{
			Name:        COMMAND_HEADLINE,
			Description: "Headline related commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        SUBCOMMAND_SUBSCRIBE,
					Description: "Subscribe headline to this channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_NEW,
							Description: "Subscribe to new headlines",
							Required:    false,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        SUBCOMMAND_UNSUBSCRIBE,
					Description: "Unsubscribe headline to this channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_NEW,
							Description: "Unsubscribe to new headlines",
							Required:    false,
						},
					},
				},
			},
		},
		{
			Name:        COMMAND_MANGA_POST,
			Description: "Manga related commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        SUBCOMMAND_SUBSCRIBE,
					Description: "Subscribe manga post to this channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_NEW,
							Description: "Subscribe to new manga posts",
							Required:    false,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        SUBCOMMAND_UNSUBSCRIBE,
					Description: "Unsubscribe manga post to this channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Name:        SUBCOMMAND_NEW,
							Description: "Unsubscribe to new manga posts",
							Required:    false,
						},
					},
				},
			},
		},
	}
)
