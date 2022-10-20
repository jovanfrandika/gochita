package dBot

import (
	"github.com/bwmarrin/discordgo"
	uBot "github.com/jovanfrandika/gochita/internal/usecase/bot"
)

type delivery struct {
	cmds    []*discordgo.ApplicationCommand
	usecase *uBot.Usecase
}

type Delivery interface {
	InitHandler()
	RunNotifier()
	RegisterCommands()
	UnregisterCommands()
}

func New(usecase *uBot.Usecase) *delivery {
	return &delivery{
		usecase: usecase,
		cmds:    []*discordgo.ApplicationCommand{},
	}
}
