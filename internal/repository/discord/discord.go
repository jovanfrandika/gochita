package rDiscord

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	SESSION_NOT_FOUND = "Session not found"
)

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

func (client *discordBotRepo) RegisterCommand(cmd *discordgo.ApplicationCommand) (ccmd *discordgo.ApplicationCommand, err error) {
	if client.dg == nil {
		return nil, errors.New(SESSION_NOT_FOUND)
	}
	ccmd, err = client.dg.ApplicationCommandCreate(client.dg.State.User.ID, "", cmd)
	if err != nil {
		return nil, err
	}
	log.Printf(LABEL_CMD_REGISTERED, ccmd.Name)

	return ccmd, nil
}

func (client *discordBotRepo) UnregisterCommand(cmd *discordgo.ApplicationCommand) (err error) {
	if client.dg == nil {
		return errors.New(SESSION_NOT_FOUND)
	}
	err = client.dg.ApplicationCommandDelete(client.dg.State.User.ID, "", cmd.ID)
	if err != nil {
		return err
	}
	log.Printf(LABEL_CMD_UNREGISTERED, cmd.Name)

	return nil
}
