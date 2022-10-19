package dLivechart

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (d *delivery) Init() {
	go func() {
		for {
			d.AddShowEpisodes()

			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			d.AddHeadlines()

			time.Sleep(60 * time.Second)
		}
	}()
}

func (d *delivery) AddShowEpisodes() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).AddShowEpisodes(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("AddShowEpisodes start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("AddShowEpisodes start: %v; cancelled: %v;", now, err))
		}
	}
}

func (d *delivery) AddHeadlines() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var err error
	ch := make(chan int)
	go func() {
		err = (*d.usecase).AddHeadlines(ctx)
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("AddHeadlines start: %v; timeout: %v;", now, ctx.Err()))
	case <-ch:
		if err != nil {
			log.Println(fmt.Sprintf("AddHeadlines start: %v; cancelled: %v;", now, err))
		}
	}
}
