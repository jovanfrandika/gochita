package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetShowEpisodesByShowId(ctx context.Context, showId string) (showEpisodes []m.DbShowEpisode, err error) {
	showEpisodes = []m.DbShowEpisode{}
	iter := r.session.Query(queryGetShowEpisodesByShowId, showId).Iter()
	var showEpisode m.DbShowEpisode
	for iter.Scan(&showEpisode.Id, &showEpisode.ShowId, &showEpisode.Num, &showEpisode.PubDate) {
		showEpisodes = append(showEpisodes, showEpisode)
	}
	if err = iter.Close(); err != nil {
		return []m.DbShowEpisode{}, err
	}

	return showEpisodes, err
}

func (r *repository) CreateShowEpisode(ctx context.Context, showId string, showEpisode m.FeedShowEpisode) (showEpisodeId string, err error) {
	uuid := gocql.TimeUUID()
	err = r.session.Query(queryCreateShowEpisode, uuid, showId, showEpisode.Num, showEpisode.PubDate).Exec()
	if err != nil {
		return "", err
	}
	showId = uuid.String()

	return showId, nil
}