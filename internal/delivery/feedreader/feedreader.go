package dFeedReader

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (d *delivery) Init() {
	go func() {
		for {
			d.addShowEpisodes()

			time.Sleep(time.Duration(d.timeCfg.AddShowsInterval) * time.Second)
		}
	}()

	go func() {
		for {
			d.addHeadlines()

			time.Sleep(time.Duration(d.timeCfg.AddHeadlinesInterval) * time.Second)
		}
	}()

	go func() {
		for {
			d.addMangaPosts()

			time.Sleep(time.Duration(d.timeCfg.AddMangasInterval) * time.Second)
		}
	}()
}

func (d *delivery) addShowEpisodes() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).AddShowEpisodes(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("addShowEpisodes start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("addShowEpisodes start: %v; cancelled: %v;", now, err))
		}
	}
}

func (d *delivery) addHeadlines() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).AddHeadlines(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("addHeadlines start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("addHeadlines start: %v; cancelled: %v;", now, err))
		}
	}
}

func (d *delivery) addMangaPosts() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.DefaultTimeout)*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).AddMangaPosts(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("addMangaPosts start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("addMangaPosts start: %v; cancelled: %v;", now, err))
		}
	}
}
