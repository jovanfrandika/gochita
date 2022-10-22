package uBot

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
)

func (u *usecase) SubscribeNewMangaPost(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_NEW_MANGA_POST, referenceId, NO_CONTEXT_ID)
	if err == gocql.ErrNotFound {
		_, err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_NEW_MANGA_POST, referenceId, NO_CONTEXT_ID)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_NEW_MANGA_POST, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_NEW_MANGA_POST_SUBSCRIPTION, err
	}

	return LABEL_SUCCESS_NEW_MANGA_POST_SUBSCRIPTION, nil
}

func (u *usecase) UnsubscribeNewMangaPost(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_NEW_MANGA_POST, referenceId, NO_CONTEXT_ID)
	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_NEW_MANGA_POST, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_NEW_MANGA_POST_UNSUBSCRIPTION, err
	}

	return LABEL_SUCCESS_NEW_MANGA_POST_UNSUBSCRIPTION, nil
}

func (u *usecase) NotifyNewMangaPosts(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestMangaPosts, err := (*u.dbRepo).GetMangaPostsByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	dbSubscriptions, err := (*u.dbRepo).GetSubscriptions(ctx, SUBSCRIPTION_TYPE_NEW_MANGA_POST, true)
	if err != nil {
		return err
	}

	for _, latestMangaPost := range dbLatestMangaPosts {
		for _, subscriber := range dbSubscriptions {
			dbSpecificMangaPostSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_MANGA_POST, subscriber.Id, latestMangaPost.Id)
			var channelSubscriptionId string
			if err == gocql.ErrNotFound {
				id, err := (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_MANGA_POST, subscriber.Id, latestMangaPost.Id)
				if err != nil {
					return err
				}
				channelSubscriptionId = id
			} else {
				channelSubscriptionId = dbSpecificMangaPostSubscription.Id
			}

			_, err = (*u.dbRepo).GetNotification(ctx, channelSubscriptionId)
			if err == gocql.ErrNotFound {
				content := fmt.Sprintf(LABEL_NEW_MANGA_POST, latestMangaPost.Title, latestMangaPost.PublishedAt.In(u.timeCfg.TimeLocation).Format(time.RFC850), latestMangaPost.Ref)
				msg := &discordgo.MessageSend{
					Content: content,
				}
				(*u.discordBotRepo).SendMsgToChannel(subscriber.ReferenceId, msg)
				err = (*u.dbRepo).CreateNotification(ctx, channelSubscriptionId)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
