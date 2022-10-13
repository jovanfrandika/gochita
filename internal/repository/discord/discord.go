package rDiscord

import "github.com/bwmarrin/discordgo"

func (client *discordBotRepo) AddHandler(handler interface{}) {
	if client.dg == nil {
		return
	}

	client.dg.AddHandler(handler)
}

func (client *discordBotRepo) Connect() (err error) {
	err = client.dg.Open()
	return err
}

func (client *discordBotRepo) Close() {
	if client.dg == nil {
		return
	}
	client.dg.Close()
}

func (client *discordBotRepo) SendMsgToChannel(channelId string, msg *discordgo.MessageSend) {
	if client.dg == nil {
		return
	}
	client.dg.ChannelMessageSendComplex(channelId, msg)
}
