package livechart

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jovanfrandika/livechart-notifier/config"
	dBot "github.com/jovanfrandika/livechart-notifier/internal/delivery/bot"
	rCassandra "github.com/jovanfrandika/livechart-notifier/internal/repository/cassandra"
	rDiscord "github.com/jovanfrandika/livechart-notifier/internal/repository/discord"
	uBot "github.com/jovanfrandika/livechart-notifier/internal/usecase/bot"
)

func main() {
	cfg := config.Init()

	discordBotRepo := rDiscord.New(cfg.Bot.Token)
	defer discordBotRepo.Close()

	dbRepo := rCassandra.New(cfg.DB.Clusters, cfg.DB.KeyspaceName)
	defer dbRepo.CloseConnection()

	u := uBot.New(&dbRepo, &discordBotRepo)
	d := dBot.New(&u)

	go d.RunNotifier()
	d.RegisterCommands()
	defer d.UnregisterCommands()
	d.InitHandler()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
