package dBot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	QUERY = "query"

	COMMAND_SHOW_LIST        = "show-list"
	COMMAND_SHOW_SUBSCRIBE   = "show-subscribe"
	COMMAND_SHOW_UNSUBSCRIBE = "show-unsubscribe"

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        COMMAND_SHOW_LIST,
			Description: "List subscribed shows of this channel",
		},
		{
			Name:        COMMAND_SHOW_SUBSCRIBE,
			Description: "Subscribe show to this channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        QUERY,
					Description: "Show title",
					Required:    true,
				},
			},
		},
		{
			Name:        COMMAND_SHOW_UNSUBSCRIBE,
			Description: "Unsubscribe show to this channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        QUERY,
					Description: "Show title",
					Required:    true,
				},
			},
		},
	}
)

func (d *delivery) RunNotifier() {
	go func() {
		for {
			now := time.Now()
			minute := now.Minute()
			if minute%15 == 0 {
				d.doNotifyLatestEpisodes()
			}

			time.Sleep(30 * time.Second)
		}
	}()
}

func (d *delivery) InitHandler() {
	(*d.usecase).AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch command := i.ApplicationCommandData().Name; command {
		case COMMAND_SHOW_LIST:
			d.getSubscriptions(s, i)
		case COMMAND_SHOW_SUBSCRIBE:
			d.subscribe(s, i)
		case COMMAND_SHOW_UNSUBSCRIBE:
			d.unsubscribe(s, i)
		default:
		}
	})
}

func (d *delivery) getSubscriptions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).GetSubscriptions(ctx, i.ChannelID)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
		}
	}

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: content,
	})
}

func (d *delivery) unsubscribe(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	showTitle := getQueryValue(i)
	if showTitle == "" {
		return
	}

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).Unsubscribe(ctx, i.ChannelID, showTitle)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
		}
	}

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: content,
	})
}

func (d *delivery) subscribe(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	showTitle := getQueryValue(i)
	if showTitle == "" {
		return
	}

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).Subscribe(ctx, i.ChannelID, showTitle)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
		}
	}

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: content,
	})
}

func (d *delivery) doNotifyLatestEpisodes() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).NotifyNewEpisodes(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("DoNotifyLatestEpisodes start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("DoNotifyLatestEpisodes handler start: %v; cancelled: %v;", now, ctx.Err()))
		}
	}
}

// TODO
func (d *delivery) RegisterCommands() {
	return
}

// TODO
func (d *delivery) UnregisterCommands() {
	return
}
