package uBot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (u *usecase) AddHandler(handler interface{}) {
	(*u.discordBotRepo).AddHandler(handler)
}

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

func (u *usecase) SubscribeAllShow(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_ALL_SHOW, referenceId, NO_CONTEXT_ID)
	log.Println(err)
	if err == gocql.ErrNotFound {
		err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_ALL_SHOW, referenceId, NO_CONTEXT_ID)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_ALL_SHOW, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_ALL_SHOW_SUBSCRIPTION, err
	}

	return LABEL_SUCCESS_ALL_SHOW_SUBSCRIPTION, nil
}

func (u *usecase) SubscribeSpecificShow(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), err
	}

	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	if err == gocql.ErrNotFound {
		err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), nil
}

func (u *usecase) SubscribeAllHeadline(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_ALL_HEADLINE, referenceId, NO_CONTEXT_ID)
	if err == gocql.ErrNotFound {
		err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_ALL_HEADLINE, referenceId, NO_CONTEXT_ID)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_ALL_HEADLINE, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_ALL_HEADLINE_SUBSCRIPTION, err
	}

	return LABEL_SUCCESS_ALL_HEADLINE_SUBSCRIPTION, nil
}

func (u *usecase) UnsubscribeAllShow(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_ALL_SHOW, referenceId, NO_CONTEXT_ID)
	if err == gocql.ErrNotFound {
		return LABEL_UNSUCCESS_ALL_SHOW_UNSUBSCRIPTION, err
	}

	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_ALL_SHOW, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_ALL_SHOW_UNSUBSCRIPTION, err
	}

	return LABEL_SUCCESS_ALL_SHOW_UNSUBSCRIPTION, nil
}

func (u *usecase) UnsubscribeSpecificShow(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), err
	}

	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SPECIFIC_SHOW_SUBSCRIPTION, showTitle), nil
}

func (u *usecase) UnsubscribeAllHeadline(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_ALL_HEADLINE, referenceId, NO_CONTEXT_ID)
	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_ALL_HEADLINE, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_ALL_HEADLINE_UNSUBSCRIPTION, err
	}

	return LABEL_SUCCESS_ALL_HEADLINE_UNSUBSCRIPTION, nil
}

func (u *usecase) NotifyNewShowEpisodes(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestShowEpisodes, err := (*u.dbRepo).GetShowEpisodesByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	dbAllShowSubscribers, err := (*u.dbRepo).GetSubscriptions(ctx, SUBSCRIPTION_TYPE_ALL_SHOW, true)
	if err != nil {
		return err
	}

	for _, latestEpisode := range dbLatestShowEpisodes {
		for _, subscriber := range dbAllShowSubscribers {
			_, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, subscriber.ReferenceId, latestEpisode.ShowId)
			if err == gocql.ErrNotFound {
				err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_SHOW, subscriber.ReferenceId, latestEpisode.ShowId)
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
			content = fmt.Sprintf(LABEL_NEW_SHOW_SERIES_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.Num, latestEpisode.PubDate.In(u.timeLocation).Format(time.RFC850))
		} else {
			content = fmt.Sprintf(LABEL_NEW_SHOW_MOVIE_EPISODE, showMap[latestEpisode.ShowId].Title, latestEpisode.PubDate.In(u.timeLocation).Format(time.RFC850))
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

func (u *usecase) NotifyNewHeadlines(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestHeadlines, err := (*u.dbRepo).GetHeadlinesByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	dbSubscriptions, err := (*u.dbRepo).GetSubscriptions(ctx, SUBSCRIPTION_TYPE_ALL_HEADLINE, true)
	if err != nil {
		return err
	}

	for _, latestHeadline := range dbLatestHeadlines {
		for _, subscriber := range dbSubscriptions {
			content := fmt.Sprintf(LABEL_NEW_HEADLINE, latestHeadline.Title, latestHeadline.PublishedAt.In(u.timeLocation).Format(time.RFC850), latestHeadline.Ref)
			msg := &discordgo.MessageSend{
				Content: content,
			}
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
