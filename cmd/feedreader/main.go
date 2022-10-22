package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jovanfrandika/gochita/config"
	dFeedReader "github.com/jovanfrandika/gochita/internal/delivery/feedreader"
	rCassandra "github.com/jovanfrandika/gochita/internal/repository/cassandra"
	rHttpcall "github.com/jovanfrandika/gochita/internal/repository/httpcall"
	uFeedReader "github.com/jovanfrandika/gochita/internal/usecase/feedreader"
)

func main() {
	cfg := config.Init()

	dbRepo := rCassandra.New(cfg.DB.Clusters, cfg.DB.KeyspaceName)
	defer dbRepo.CloseConnection()

	livechartClient := rHttpcall.New(&cfg.LiveChart, &cfg.Reddit, &cfg.Time)
	u := uFeedReader.New(&dbRepo, &livechartClient, &cfg.Time)
	d := dFeedReader.New(&u, &cfg.Time)

	d.Init()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
