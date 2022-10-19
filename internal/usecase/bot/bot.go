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

func (u *usecase) GetShowSubscriptions(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscriptions, err := (*u.dbRepo).GetShowSubscriptionsByReferenceId(ctx, referenceId, true)
	if err != nil {
		return DEFAULT_ERROR, err
	}

	if len(dbSubscriptions) <= 0 {
		return NO_SUBSCRIPTIONS, nil
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

func (u *usecase) SubscribeShow(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SHOW_UNSUBSCRIPTION, showTitle), err
	}

	dbSubscription, err := (*u.dbRepo).GetShowSubscription(ctx, referenceId, dbShow.Id)
	if err == gocql.ErrNotFound || !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleShowSubscription(ctx, true, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SHOW_SUBSCRIPTION, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SHOW_SUBSCRIPTION, showTitle), nil
}

func (u *usecase) SubscribeHeadline(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, err := (*u.dbRepo).GetHeadlineSubscription(ctx, referenceId)
	if err == gocql.ErrNotFound || !dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleHeadlineSubscription(ctx, true, referenceId)
	}
	if err != nil {
		return LABEL_UNSUCCESS_HEADLINE_SUBSCRIPTION, err
	}

	return LABEL_SUCCESS_HEADLINE_SUBSCRIPTION, nil
}

func (u *usecase) UnsubscribeShow(ctx context.Context, referenceId string, showTitle string) (content string, err error) {
	dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showTitle)
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SHOW_UNSUBSCRIPTION, showTitle), err
	}

	dbSubscription, _ := (*u.dbRepo).GetShowSubscription(ctx, referenceId, dbShow.Id)
	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleShowSubscription(ctx, false, referenceId, dbShow.Id)
	}
	if err != nil {
		return fmt.Sprintf(LABEL_UNSUCCESS_SHOW_UNSUBSCRIPTION, showTitle), err
	}

	return fmt.Sprintf(LABEL_SUCCESS_SHOW_UNSUBSCRIPTION, showTitle), nil
}

func (u *usecase) UnsubscribeHeadline(ctx context.Context, referenceId string) (content string, err error) {
	dbSubscription, _ := (*u.dbRepo).GetHeadlineSubscription(ctx, referenceId)
	if err == gocql.ErrNotFound {
		return LABEL_UNSUCCESS_HEADLINE_UNSUBSCRIPTION, err
	}

	if dbSubscription.IsEnabled {
		err = (*u.dbRepo).ToggleHeadlineSubscription(ctx, false, referenceId)
	}
	if err != nil {
		return LABEL_UNSUCCESS_HEADLINE_UNSUBSCRIPTION, err
	}

	return LABEL_SUCCESS_HEADLINE_UNSUBSCRIPTION, nil
}

func (u *usecase) NotifyNewHeadlines(ctx context.Context) (err error) {
	now := time.Now()
	dbLatestHeadlines, err := (*u.dbRepo).GetHeadlinesByRange(ctx, now.AddDate(0, 0, -1), now)
	if err != nil {
		return err
	}

	dbHeadlineSubscriptions, err := (*u.dbRepo).GetHeadlineSubscriptions(ctx)
	if err != nil {
		return err
	}

	for _, latestHeadline := range dbLatestHeadlines {
		for _, subscriber := range dbHeadlineSubscriptions {
			content := fmt.Sprintf(LABEL_NEW_HEADLINE, latestHeadline.Title, latestHeadline.PublishedAt.In(u.timeLocation).Format(time.RFC850), latestHeadline.Ref)
			msg := &discordgo.MessageSend{
				Content: content,
			}
			_, err := (*u.dbRepo).GetHeadlineNotification(ctx, subscriber.ReferenceId, latestHeadline.Id)
			if err == gocql.ErrNotFound {
				(*u.discordBotRepo).SendMsgToChannel(subscriber.ReferenceId, msg)
				err = (*u.dbRepo).CreateHeadlineNotification(ctx, subscriber.ReferenceId, latestHeadline.Id)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
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
			dbChannelShowSubscriptions, err := (*u.dbRepo).GetShowSubscriptionsByShowId(ctx, latestEpisode.ShowId, true)
			if err != nil {
				return err
			}
			showIdToChannelMap[latestEpisode.ShowId] = dbChannelShowSubscriptions
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
		for _, channel := range showIdToChannelMap[latestEpisode.ShowId] {
			_, err := (*u.dbRepo).GetShowNotification(ctx, channel.ReferenceId, latestEpisode.ShowId)
			if err == gocql.ErrNotFound {
				(*u.discordBotRepo).SendMsgToChannel(channel.ReferenceId, msg)
				err = (*u.dbRepo).CreateShowNotification(ctx, channel.ReferenceId, latestEpisode.ShowId)
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
