package dBot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (d *delivery) RunNotifier() {
	go func() {
		for {
			now := time.Now()
			minute := now.Minute()
			if minute%5 == 0 {
				d.doNotifyLatestEpisodes()
			}

			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			now := time.Now()
			minute := now.Minute()
			if minute%2 == 0 {
				d.doNotifyLatestHeadlines()
			}

			time.Sleep(10 * time.Second)
		}
	}()
}

func (d *delivery) InitHandler() {
	(*d.usecase).AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch command := i.ApplicationCommandData().Name; command {
		case COMMAND_SHOW_LIST:
			d.getShowSubscriptions(s, i)
		case COMMAND_SHOW_SUBSCRIBE_ALL:
			d.subscribeAllShow(s, i)
		case COMMAND_SHOW_UNSUBSCRIBE_ALL:
			d.unsubscribeAllShow(s, i)
		case COMMAND_SHOW_SUBSCRIBE:
			d.subscribeSpecificShow(s, i)
		case COMMAND_SHOW_UNSUBSCRIBE:
			d.unsubscribeSpecificShow(s, i)
		case COMMAND_HEADLINE_SUBSCRIBE_ALL:
			d.subscribeAllHeadline(s, i)
		case COMMAND_HEADLINE_UNSUBSCRIBE_ALL:
			d.unsubscribeAllHeadline(s, i)
		default:
		}
	})
}

func (d *delivery) subscribeAllShow(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).SubscribeAllShow(ctx, i.ChannelID)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) unsubscribeAllShow(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).UnsubscribeAllShow(ctx, i.ChannelID)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) getShowSubscriptions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).GetShowSubscriptions(ctx, i.ChannelID)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) subscribeSpecificShow(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		content, err = (*d.usecase).SubscribeSpecificShow(ctx, i.ChannelID, showTitle)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) unsubscribeSpecificShow(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		content, err = (*d.usecase).UnsubscribeSpecificShow(ctx, i.ChannelID, showTitle)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) subscribeAllHeadline(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).SubscribeAllHeadline(ctx, i.ChannelID)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) unsubscribeAllHeadline(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).UnsubscribeAllHeadline(ctx, i.ChannelID)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("%v handler start: %v; timeout: %v;", i.ApplicationCommandData().Name, now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("%v handler start: %v; cancelled: %v;", i.ApplicationCommandData().Name, now, err))
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (d *delivery) doNotifyLatestEpisodes() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).NotifyNewShowEpisodes(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("DoNotifyLatestEpisodes start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("DoNotifyLatestEpisodes handler start: %v; cancelled: %v;", now, err))
		}
	}
}

func (d *delivery) doNotifyLatestHeadlines() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).NotifyNewHeadlines(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("NotifyNewHeadlines start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("NotifyNewHeadlines handler start: %v; cancelled: %v;", now, err))
		}
	}
}

func (d *delivery) RegisterCommands() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		d.cmds, err = (*d.usecase).RegisterCommands(ctx, commands)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Fatal(fmt.Sprintf("RegisterCommands timeout: %v;", ctx.Err()))
	case <-ch:
		if err != nil {
			log.Fatal(fmt.Sprintf("RegisterCommands cancelled: %v;", err))
		}
	}
}

func (d *delivery) UnregisterCommands() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).UnregisterCommands(ctx, d.cmds)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("UnregisterCommands timeout: %v;", ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("UnregisterCommands cancelled: %v;", err))
		}
	}
}
