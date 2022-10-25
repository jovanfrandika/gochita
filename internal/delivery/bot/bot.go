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
			if minute%d.timeCfg.NotifyShowsInterval == 0 {
				d.doNotifyLatestEpisodes()
			}

			time.Sleep(20 * time.Second)
		}
	}()

	go func() {
		for {
			now := time.Now()
			minute := now.Minute()
			if minute%d.timeCfg.NotifyHeadlinesInterval == 0 {
				d.doNotifyLatestHeadlines()
			}

			time.Sleep(20 * time.Second)
		}
	}()

	go func() {
		for {
			now := time.Now()
			minute := now.Minute()
			if minute%d.timeCfg.NotifyMangasInterval == 0 {
				d.doNotifyLatestMangaPosts()
			}

			time.Sleep(20 * time.Second)
		}
	}()
}

func (d *delivery) InitHandler() {
	(*d.usecase).AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch command := i.ApplicationCommandData().Name; command {
		case COMMAND_SHOW:
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case SUBCOMMAND_LIST:
				d.getShowSubscriptions(s, i)
			case SUBCOMMAND_SUBSCRIBE:
				options = options[0].Options
				switch options[0].Name {
				case SUBCOMMAND_NEW:
					d.subscribeNewShow(s, i)
				case SUBCOMMAND_ONE:
					d.subscribeSpecificShow(s, i)
				default:
				}
			case SUBCOMMAND_UNSUBSCRIBE:
				options = options[0].Options
				switch options[0].Name {
				case SUBCOMMAND_ALL:
					d.unsubscribeAllShow(s, i)
				case SUBCOMMAND_NEW:
					d.unsubscribeNewShow(s, i)
				case SUBCOMMAND_ONE:
					d.unsubscribeSpecificShow(s, i)
				default:
				}
			default:
			}
		case COMMAND_HEADLINE:
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case SUBCOMMAND_SUBSCRIBE:
				options = options[0].Options

				switch options[0].Name {
				case SUBCOMMAND_NEW:
					d.subscribeNewHeadline(s, i)
				default:
				}
			case SUBCOMMAND_UNSUBSCRIBE:
				options = options[0].Options

				switch options[0].Name {
				case SUBCOMMAND_NEW:
					d.unsubscribeNewHeadline(s, i)
				default:
				}
			default:
			}
		case COMMAND_MANGA_POST:
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case SUBCOMMAND_SUBSCRIBE:
				options = options[0].Options

				switch options[0].Name {
				case SUBCOMMAND_NEW:
					d.subscribeNewMangaPost(s, i)
				default:
				}
			case SUBCOMMAND_UNSUBSCRIBE:
				options = options[0].Options

				switch options[0].Name {
				case SUBCOMMAND_NEW:
					d.unsubscribeNewMangaPost(s, i)
				default:
				}
			default:
			}
		default:
		}
	})
}

func (d *delivery) RegisterCommands() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.SetCommandsTimeout)*time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.timeCfg.SetCommandsTimeout)*time.Second)
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
