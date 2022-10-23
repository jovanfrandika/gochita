package dBot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (d *delivery) getShowSubscriptions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
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

func (d *delivery) subscribeNewShow(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).SubscribeNewShow(ctx, i.ChannelID)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
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

func (d *delivery) unsubscribeNewShow(s *discordgo.Session, i *discordgo.InteractionCreate) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
	defer cancel()

	var content string
	var err error
	ch := make(chan int)
	go func() {
		content, err = (*d.usecase).UnsubscribeNewShow(ctx, i.ChannelID)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
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

func (d *delivery) doNotifyLatestEpisodes() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.NotifyTimeout)*time.Second)
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
