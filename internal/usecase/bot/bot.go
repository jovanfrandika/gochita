package uBot

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

const (
	LABEL_NEW_EPISODE        = "***Reminder!!!***\nTitle: %v\nEpisode: %v\nPublished Date: %v\n"
	LABEL_SUBSCRIPTION_TITLE = "%v. %v\n"

	LABEL_SUCCESS_SUBSCRIBED   = "%v successfully subscribed!"
	LABEL_UNSUCCESS_SUBSCRIBED = "%v subscription failed :("

	LABEL_SUCCESS_UNSUBSCRIBED   = "%v successfully unsubscribed!"
	LABEL_UNSUCCESS_UNSUBSCRIBED = "%v unsubscription failed :("

	NO_SUBSCRIPTIONS = "No subscriptions"
	DEFAULT_ERROR    = "Oops something went wrong!"
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

	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, dbShow.Id, referenceId)
	if err == gocql.ErrNotFound {
		err = (*u.dbRepo).CreateSubscription(ctx, dbShow.Id, referenceId)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, dbShow.Id, referenceId)
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

	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, dbShow.Id, referenceId)
	if err != nil && dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, dbShow.Id, referenceId)
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

		msg := &discordgo.MessageSend{
			Content: fmt.Sprintf(LABEL_NEW_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.Num, latestEpisode.PubDate.Format(time.RFC850)),
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
