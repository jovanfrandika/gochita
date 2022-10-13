package uLivechart

import (
	"context"

	"github.com/gocql/gocql"
)

func (u *usecase) AddShowEpisodes(ctx context.Context) (err error) {
	showMap, err := (*u.client).GetLatestEpisodes()
	if err != nil {
		return err
	}

	for showName, showDetail := range showMap {
		var showId string
		dbShow, err := (*u.dbRepo).GetShowByTitle(ctx, showName)
		if err != nil {
			if err != gocql.ErrNotFound {
				return err
			}
			showId, err = (*u.dbRepo).CreateShow(ctx, showDetail)
			if err != nil {
				return err
			}
		} else {
			showId = dbShow.Id
		}

		dbShowEpisodes, err := (*u.dbRepo).GetShowEpisodesByShowId(ctx, showId)
		if err != nil {
			return err
		}

	EpisodeLoop:
		for showEpisodeNum, showEpisodeDetail := range showDetail.ShowEpisodeMap {
			for _, dbEpisode := range dbShowEpisodes {
				if dbEpisode.Num == showEpisodeNum {
					continue EpisodeLoop
				}
			}

			_, err = (*u.dbRepo).CreateShowEpisode(ctx, showId, showEpisodeDetail)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
