package model

import "time"

type DbChannelShowEpisodeNotification struct {
	ShowEpisodeId string
	ReferenceId   string
	NotifiedAt    time.Time
}
