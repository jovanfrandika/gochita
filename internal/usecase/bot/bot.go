package uBot

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (u *usecase) AddHandler(handler interface{}) {
	(*u.discordBotRepo).AddHandler(handler)
}

func (u *usecase) GetSubscriptions(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscriptions, err := (*u.dbRepo).GetSubscriptionsByReferenceId(ctx, referenceId)
	if err != nil {
		return DEFAULT_ERROR, err
	}

	if len(dbSubscriptions) < 0 {
		(*u.discordBotRepo).SendMsgToChannel(referenceId, &discordgo.MessageSend{
			Content: NO_SUBSCRIPTIONS,
		})
		return DEFAULT_ERROR, err
	}

	content = "Subscriptions\n"
	for idx, subscription := range dbSubscriptions {
		dbShow, err := (*u.dbRepo).GetShowById(ctx, subscription.ShowId)
		if err != nil {
			return DEFAULT_ERROR, err
		}
		content += fmt.Sprintf(LABEL_SUBSCRIPTION_TITLE, idx+1, dbShow.Title)
	}

	return content, nil
}

func (u *usecase) Subscribe(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SUBSCRIBED, showTitle), err
	}

	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, referenceId, dbShow.Id)
	if err == gocql.ErrNotFound {
		err = (*u.dbRepo).CreateSubscription(ctx, referenceId, dbShow.Id)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SUBSCRIBED, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SUBSCRIBED, showTitle), nil
}

func (u *usecase) Unsubscribe(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_UNSUBSCRIBED, showTitle), err
	}

	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, referenceId, dbShow.Id)
	if err != nil && dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_UNSUBSCRIBED, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_UNSUBSCRIBED, showTitle), nil
}

func (u *usecase) NotifyNewEpisodes(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestShowEpisodes, err := (*u.dbRepo).GetShowEpisodesByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	showMap := map[string]m.DbShow{}
	showIdToChannelMap := map[string][]m.DbChannelShowSubscription{}
	for _, latestEpisode := range dbLatestShowEpisodes {
		if _, exists := showMap[latestEpisode.ShowId]; !exists {
			dbShow, err := (*u.dbRepo).GetShowById(ctx, latestEpisode.ShowId)
			if err != nil {
				return err
			}
			showMap[latestEpisode.ShowId] = dbShow
		}

		if _, exists := showIdToChannelMap[latestEpisode.ShowId]; !exists {
			dbChannelShowSubscriptions, err := (*u.dbRepo).GetSubscriptionsByShowId(ctx, latestEpisode.ShowId)
			if err != nil {
				return err
			}
			showIdToChannelMap[latestEpisode.ShowId] = dbChannelShowSubscriptions
		}

		var content string
		if latestEpisode.Num != 0 {
			content = fmt.Sprintf(LABEL_NEW_SERIES_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.Num, latestEpisode.PubDate.Format(time.RFC850))
		} else {
			content = fmt.Sprintf(LABEL_NEW_MOVIE_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.PubDate.Format(time.RFC850))
		}
		msg := &discordgo.MessageSend{
			Content: content,
			Embeds: []*discordgo.MessageEmbed{
				{
					Image: &discordgo.MessageEmbedImage{
						URL: showMap[latestEpisode.ShowId].Thumbnail,
					},
					Type: discordgo.EmbedTypeImage,
				},
			},
		}
		for _, channel := range showIdToChannelMap[latestEpisode.ShowId] {
			_, err := (*u.dbRepo).GetNotification(ctx, latestEpisode.ShowId, channel.ReferenceId)
			if err == gocql.ErrNotFound {
				(*u.discordBotRepo).SendMsgToChannel(channel.ReferenceId, msg)
				err = (*u.dbRepo).CreateNotification(ctx, latestEpisode.ShowId, channel.ReferenceId)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (u *usecase) RegisterCommands(ctx context.Context, cmds []*discordgo.ApplicationCommand) (ccmds []*discordgo.ApplicationCommand, err error) {
	ccmds = make([]*discordgo.ApplicationCommand, len(cmds))
	for i, cmd := range cmds {
		ccmd, err := (*u.discordBotRepo).RegisterCommand(cmd)
		if err != nil {
			return []*discordgo.ApplicationCommand{}, err
		}
		ccmds[i] = ccmd
	}

	return ccmds, nil
}

func (u *usecase) UnregisterCommands(ctx context.Context, cmds []*discordgo.ApplicationCommand) (err error) {
	for _, cmd := range cmds {
		err = (*u.discordBotRepo).UnregisterCommand(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
