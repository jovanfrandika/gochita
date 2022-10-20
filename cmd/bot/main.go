package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jovanfrandika/gochita/config"
	dBot "github.com/jovanfrandika/gochita/internal/delivery/bot"
	rCassandra "github.com/jovanfrandika/gochita/internal/repository/cassandra"
	rDiscord "github.com/jovanfrandika/gochita/internal/repository/discord"
	uBot "github.com/jovanfrandika/gochita/internal/usecase/bot"
)

func main() {
	cfg := config.Init()

	timeLocation, err := time.LoadLocation(cfg.Time.Timezone)
	if err != nil {
		log.Fatal(err.Error())
	}

	discordBotRepo := rDiscord.New(cfg.Bot.Token)
	if err := discordBotRepo.Connect(); err != nil {
		log.Fatal(err)
	}
	defer discordBotRepo.Close()

	dbRepo := rCassandra.New(cfg.DB.Clusters, cfg.DB.KeyspaceName)
	defer dbRepo.CloseConnection()

	u := uBot.New(&dbRepo, &discordBotRepo, timeLocation)
	d := dBot.New(&u)

	d.RegisterCommands()
	d.InitHandler()
	defer d.UnregisterCommands()

	go d.RunNotifier()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
