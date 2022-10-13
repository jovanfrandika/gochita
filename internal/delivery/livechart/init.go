package dLivechart

import uLivechart "github.com/jovanfrandika/livechart-notifier/internal/usecase/livechart"

type delivery struct {
	usecase *uLivechart.Usecase
}

type Delivery interface {
	AddShowEpisodes()
}

func New(usecase *uLivechart.Usecase) *delivery {
	return &delivery{
		usecase: usecase,
	}
}
