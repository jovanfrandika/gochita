package dBot

import uBot "github.com/jovanfrandika/livechart-notifier/internal/usecase/bot"

type delivery struct {
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
	}
}
