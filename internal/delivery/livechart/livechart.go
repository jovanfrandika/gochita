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
			log.Println(fmt.Sprintf("AddShowEpisodes start: %v; cancelled: %v;", now, ctx.Err()))
		}
	}
}
