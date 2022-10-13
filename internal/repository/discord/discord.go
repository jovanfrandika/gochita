package rDiscord

import (
	"fmt"
)

const (
	LABEL_NEW_EPISODE = "***Reminder!!!***\nTitle: %v\nEpisode: %v\nPublished Date: %v\n"
)

func (client *BotClient) Connect() error {
	tries := 0

	var err error
	for tries < 3 {
		err = client.dg.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
		}
		tries++
	}

	return err
}

func (client *BotClient) Close() {
	if client.dg == nil {
		return
	}
	client.dg.Close()
}
