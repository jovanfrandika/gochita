package dFeedReader

import (
	m "github.com/jovanfrandika/gochita/domain"
	uFeedReader "github.com/jovanfrandika/gochita/internal/usecase/feedreader"
)

type delivery struct {
	usecase *uFeedReader.Usecase
	timeCfg *m.TimeConfig
}

type Delivery interface {
	Init()
}

func New(usecase *uFeedReader.Usecase, timeCfg *m.TimeConfig) *delivery {
	return &delivery{
		usecase: usecase,
		timeCfg: timeCfg,
	}
}
