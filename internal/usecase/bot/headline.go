package uBot

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
)

func (u *usecase) SubscribeNewHeadline(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_NEW_HEADLINE, referenceId, NO_CONTEXT_ID)
	if err == gocql.ErrNotFound {
		_, err = (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_NEW_HEADLINE, referenceId, NO_CONTEXT_ID)
	} else if !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, true, SUBSCRIPTION_TYPE_NEW_HEADLINE, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_NEW_HEADLINE_SUBSCRIPTION, err
	}

	return LABEL_SUCCESS_NEW_HEADLINE_SUBSCRIPTION, nil
}

func (u *usecase) UnsubscribeNewHeadline(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, _ := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_NEW_HEADLINE, referenceId, NO_CONTEXT_ID)
	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleSubscription(ctx, false, SUBSCRIPTION_TYPE_NEW_HEADLINE, referenceId, NO_CONTEXT_ID)
	}
	if err != nil {
		return LABEL_UNSUCCESS_NEW_HEADLINE_UNSUBSCRIPTION, err
	}

	return LABEL_SUCCESS_NEW_HEADLINE_UNSUBSCRIPTION, nil
}

func (u *usecase) NotifyNewHeadlines(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestHeadlines, err := (*u.dbRepo).GetHeadlinesByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	dbSubscriptions, err := (*u.dbRepo).GetSubscriptions(ctx, SUBSCRIPTION_TYPE_NEW_HEADLINE, true)
	if err != nil {
		return err
	}

	for _, latestHeadline := range dbLatestHeadlines {
		for _, subscriber := range dbSubscriptions {
			dbSpecificHeadlineSubscription, err := (*u.dbRepo).GetSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_HEADLINE, subscriber.Id, latestHeadline.Id)
			var channelSubscriptionId string
			if err == gocql.ErrNotFound {
				id, err := (*u.dbRepo).CreateSubscription(ctx, SUBSCRIPTION_TYPE_SPECIFIC_HEADLINE, subscriber.Id, latestHeadline.Id)
				if err != nil {
					return err
				}
				channelSubscriptionId = id
			} else {
				channelSubscriptionId = dbSpecificHeadlineSubscription.Id
			}

			_, err = (*u.dbRepo).GetNotification(ctx, channelSubscriptionId)
			if err == gocql.ErrNotFound {
				content := fmt.Sprintf(LABEL_NEW_HEADLINE, latestHeadline.Title, latestHeadline.PublishedAt.In(u.timeCfg.TimeLocation).Format(time.RFC850), latestHeadline.Ref)
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
