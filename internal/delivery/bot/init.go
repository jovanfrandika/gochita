package dBot

import (
	"github.com/bwmarrin/discordgo"
	m "github.com/jovanfrandika/gochita/domain"
	uBot "github.com/jovanfrandika/gochita/internal/usecase/bot"
)

type delivery struct {
	cmds    []*discordgo.ApplicationCommand
	usecase *uBot.Usecase
	timeCfg *m.TimeConfig
}

type Delivery interface {
	InitHandler()
	RunNotifier()
	RegisterCommands()
	UnregisterCommands()
}

func New(usecase *uBot.Usecase, timeCfg *m.TimeConfig) *delivery {
	return &delivery{
		usecase: usecase,
		timeCfg: timeCfg,
		cmds:    []*discordgo.ApplicationCommand{},
	}
}
