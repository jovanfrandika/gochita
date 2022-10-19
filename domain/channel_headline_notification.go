package model

import "time"

type DbChannelHeadlineNotification struct {
	HeadlineId  string
	ReferenceId string
	NotifiedAt  time.Time
}
