package uBot

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (u *usecase) AddHandler(handler interface{}) {
	(*u.discordBotRepo).AddHandler(handler)
}

func (u *usecase) RegisterCommands(ctx context.Context, cmds []*discordgo.ApplicationCommand) (ccmds []*discordgo.ApplicationCommand, err error) {
	ccmds = make([]*discordgo.ApplicationCommand, len(cmds))
	for i, cmd := range cmds {
		ccmd, err := (*u.discordBotRepo).RegisterCommand(cmd)
		if err != nil {
			return []*discordgo.ApplicationCommand{}, err
		}
		ccmds[i] = ccmd
	}

	return ccmds, nil
}

func (u *usecase) UnregisterCommands(ctx context.Context, cmds []*discordgo.ApplicationCommand) (err error) {
	for _, cmd := range cmds {
		err = (*u.discordBotRepo).UnregisterCommand(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
