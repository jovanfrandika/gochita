package dLivechart

import uLivechart "github.com/jovanfrandika/gochita/internal/usecase/livechart"

type delivery struct {
	usecase *uLivechart.Usecase
}

type Delivery interface {
	Init()
	AddShowEpisodes()
}

func New(usecase *uLivechart.Usecase) *delivery {
	return &delivery{
		usecase: usecase,
	}
}
