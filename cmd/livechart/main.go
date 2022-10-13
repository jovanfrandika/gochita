package livechart

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jovanfrandika/livechart-notifier/config"
	dLivechart "github.com/jovanfrandika/livechart-notifier/internal/delivery/livechart"
	rCassandra "github.com/jovanfrandika/livechart-notifier/internal/repository/cassandra"
	rHttpcall "github.com/jovanfrandika/livechart-notifier/internal/repository/httpcall"
	uLivechart "github.com/jovanfrandika/livechart-notifier/internal/usecase/livechart"
)

func main() {
	cfg := config.Init()

	timeLocation, err := time.LoadLocation(cfg.Time.Timezone)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbRepo := rCassandra.New(cfg.DB.Clusters, cfg.DB.KeyspaceName)
	defer dbRepo.CloseConnection()

	livechartClient := rHttpcall.New(cfg.LiveChart.BaseUrl, timeLocation)
	u := uLivechart.New(&dbRepo, &livechartClient)
	d := dLivechart.New(&u)

	d.Init()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
