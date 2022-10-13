package dBot

import "github.com/bwmarrin/discordgo"

func getQueryValue(i *discordgo.InteractionCreate) string {
	options := i.ApplicationCommandData().Options
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
