package dBot

import (
	"github.com/bwmarrin/discordgo"
)

func flattenOptions(options []*discordgo.ApplicationCommandInteractionDataOption) []*discordgo.ApplicationCommandInteractionDataOption {
	for i := 0; i < len(options); i++ {
		if len(options[i].Options) > 0 {
			newOptions := flattenOptions(options[i].Options)
			options = append(options, newOptions...)
		}
	}

	return options
}

func getQueryValue(i *discordgo.InteractionCreate) string {
	options := flattenOptions(i.ApplicationCommandData().Options)
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	stringOption, exists := optionMap[QUERY]
	if !exists {
		return ""
	}

	return stringOption.StringValue()
}
