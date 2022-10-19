package model

import "time"

type DbChannelNotification struct {
	ChannelSubscriptionId string
	NotifiedAt            time.Time
}
