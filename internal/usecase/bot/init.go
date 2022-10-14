package uBot

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	rCassandra "github.com/jovanfrandika/livechart-notifier/internal/repository/cassandra"
	rDiscord "github.com/jovanfrandika/livechart-notifier/internal/repository/discord"
)

type usecase struct {
	dbRepo         *rCassandra.Repository
	discordBotRepo *rDiscord.DiscordBotRepo
	timeLocation   *time.Location
}

type Usecase interface {
	AddHandler(handler interface{})
	GetSubscriptions(ctx context.Context, referenceId string) (content string, err error)
	Subscribe(ctx context.Context, referenceId string, showTitle string) (content string, err error)
	Unsubscribe(ctx context.Context, referenceId string, showTitle string) (content string, err error)
	NotifyNewEpisodes(ctx context.Context) (err error)
	RegisterCommands(ctx context.Context, cmds []*discordgo.ApplicationCommand) (ccmds []*discordgo.ApplicationCommand, err error)
	UnregisterCommands(ctx context.Context, cmds []*discordgo.ApplicationCommand) (err error)
}

func New(dbRepo *rCassandra.Repository, discordBotRepo *rDiscord.DiscordBotRepo, timeLocation *time.Location) Usecase {
	return &usecase{
		dbRepo:         dbRepo,
		discordBotRepo: discordBotRepo,
		timeLocation:   timeLocation,
	}
}
