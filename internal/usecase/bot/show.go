package uBot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/gochita/domain"
)

func (u *usecase) GetShowSubscriptions(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscriptions, err := (*u.dbRepo).GetSubscriptionsByReferenceId(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, true)
	if err != nil {
		return DEFAULT_ERROR, err
	}

	if len(dbSubscriptions) <= 0 {
		return NO_SUBSCRIPTIONS, nil
	}

	content = "Subscriptions\n"
	for idx, subscription := range dbSubscriptions {
		dbShow, err := (*u.dbRepo).GetShowById(ctx, subscription.ContextId)
		if err != nil {
			return DEFAULT_ERROR, err
		}
		content += fmt.Sprintf(LABEL_SUBSCRIPTION_TITLE, idx+1, dbShow.Title)
	}

	return content, nil
}

func (u *usecase) SubscribeNewShow(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_NEW_SHOW, referenceId, NO_CONTEXT_ID)
	log.Println(err)
	if err == gocql.ErrNotFound {
		_, err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_NEW_SHOW, referenceId, NO_CONTEXT_ID)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_NEW_SHOW, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_NEW_SHOW_SUBSCRIPTION, err
	}

	return LABEL_SUCCESS_NEW_SHOW_SUBSCRIPTION, nil
}

func (u *usecase) SubscribeSpecificShow(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), err
	}

	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	if err == gocql.ErrNotFound {
		_, err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), nil
}

func (u *usecase) UnsubscribeAllShow(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscriptions, err := (*u.dbRepo).GetSubscriptionsByReferenceId(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, true)
	if err != nil {
		return DEFAULT_ERROR, err
	}

	if len(dbSubscriptions) <= 0 {
		return NO_SUBSCRIPTIONS, nil
	}

	contextIds := []string{}
	for _, subscription := range dbSubscriptions {
		contextIds = append(contextIds, subscription.ContextId)
	}
	err = (*u.dbRepo).ToggleSubscriptions(ctx, false, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, contextIds)
	if err != nil {
		return DEFAULT_ERROR, err
	}

	return DEFAULT_SUCCESS, nil
}

func (u *usecase) UnsubscribeNewShow(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_NEW_SHOW, referenceId, NO_CONTEXT_ID)
	if err == gocql.ErrNotFound {
		return LABEL_UNSUCCESS_NEW_SHOW_UNSUBSCRIPTION, err
	}

	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_NEW_SHOW, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_NEW_SHOW_UNSUBSCRIPTION, err
	}

	return LABEL_SUCCESS_NEW_SHOW_UNSUBSCRIPTION, nil
}

func (u *usecase) UnsubscribeSpecificShow(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_UNSUBSCRIPTION, showTitle), err
	}

	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_UNSUBSCRIPTION, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SPECIFIC_SHOW_UNSUBSCRIPTION, showTitle), nil
}

func (u *usecase) NotifyNewShowEpisodes(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestShowEpisodes, err := (*u.dbRepo).GetShowEpisodesByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	dbNewShowSubscribers, err := (*u.dbRepo).GetSubscriptions(ctx, SUBSCRIPTION_TYPE_NEW_SHOW, true)
	if err != nil {
		return err
	}

	for _, latestEpisode := range dbLatestShowEpisodes {
		for _, subscriber := range dbNewShowSubscribers {
			_, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, subscriber.ReferenceId, latestEpisode.ShowId)
			if err == gocql.ErrNotFound {
				_, err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, subscriber.ReferenceId, latestEpisode.ShowId)
			}
			if err != nil {
				return err
			}
		}
	}

	showMap := map[string]m.DbShow{}
	showIdToSubscriptionMap := map[string][]m.DbChannelSubscription{}
	for _, latestEpisode := range dbLatestShowEpisodes {
		if _, exists := showMap[latestEpisode.ShowId]; !exists {
			dbShow, err := (*u.dbRepo).GetShowById(ctx, latestEpisode.ShowId)
			if err != nil {
				return err
			}
			showMap[latestEpisode.ShowId] = dbShow
		}

		if _, exists := showIdToSubscriptionMap[latestEpisode.ShowId]; !exists {
			dbChannelSubscriptions, err := (*u.dbRepo).GetSubscriptionsByContextId(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, latestEpisode.ShowId, true)
			if err != nil {
				return err
			}
			showIdToSubscriptionMap[latestEpisode.ShowId] = dbChannelSubscriptions
		}

		var content string
		if latestEpisode.Num != 0 {
			content = fmt.Sprintf(LABEL_NEW_SHOW_SERIES_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.Num, latestEpisode.PubDate.In(u.timeCfg.TimeLocation).Format(time.RFC850))
		} else {
			content = fmt.Sprintf(LABEL_NEW_SHOW_MOVIE_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.PubDate.In(u.timeCfg.TimeLocation).Format(time.RFC850))
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
		for _, subscriber := range showIdToSubscriptionMap[latestEpisode.ShowId] {
			_, err := (*u.dbRepo).GetNotification(ctx, subscriber.Id)
			if err == gocql.ErrNotFound {
				(*u.discordBotRepo).SendMsgToChannel(subscriber.ReferenceId, msg)
				err = (*u.dbRepo).CreateNotification(ctx, subscriber.Id)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
