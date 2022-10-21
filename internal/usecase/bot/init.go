package uBot

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	rCassandra "github.com/jovanfrandika/gochita/internal/repository/cassandra"
	rDiscord "github.com/jovanfrandika/gochita/internal/repository/discord"
)

type usecase struct {
	dbRepo         *rCassandra.Repository
	discordBotRepo *rDiscord.DiscordBotRepo
	timeLocation   *time.Location
}

type Usecase interface {
	AddHandler(handler interface{})

	GetShowSubscriptions(ctx context.Context, referenceId string) (content string, err error)

	SubscribeNewShow(ctx context.Context, referenceId string) (content string, err error)
	SubscribeSpecificShow(ctx context.Context, referenceId string, showTitle string) (content string, err error)
	SubscribeNewHeadline(ctx context.Context, referenceId string) (content string, err error)

	UnsubscribeNewShow(ctx context.Context, referenceId string) (content string, err error)
	UnsubscribeSpecificShow(ctx context.Context, referenceId string, showTitle string) (content string, err error)
	UnsubscribeNewHeadline(ctx context.Context, referenceId string) (content string, err error)

	NotifyNewShowEpisodes(ctx context.Context) (err error)
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
