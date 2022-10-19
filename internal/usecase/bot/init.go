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

	GetShowSubscriptions(ctx context.Context, referenceId string) (content string, err error)

	SubscribeShow(ctx context.Context, referenceId string, showTitle string) (content string, err error)
	SubscribeHeadline(ctx context.Context, referenceId string) (content string, err error)

	UnsubscribeShow(ctx context.Context, referenceId string, showTitle string) (content string, err error)
	UnsubscribeHeadline(ctx context.Context, referenceId string) (content string, err error)

	NotifyNewEpisodes(ctx context.Context) (err error)
	NotifyNewHeadlines(ctx context.Context) (err error)

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
