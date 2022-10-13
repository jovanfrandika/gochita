package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetShowByTitle(ctx context.Context, title string) (show m.DbShow, err error) {
	show = m.DbShow{}
	err = r.session.Query(queryGetShowByTitle, title).Consistency(gocql.One).Scan(&show.Id, &show.Title, &show.Thumbnail, &show.Category, &show.Ref)
	if err != nil {
		return m.DbShow{}, err
	}

	return show, err
}

func (r *repository) CreateShow(ctx context.Context, show m.FeedShow) (showId string, err error) {
	uuid := gocql.TimeUUID()
	err = r.session.Query(queryCreateShow, uuid, show.Title, show.Thumbnail, show.Category, show.Ref).Exec()
	if err != nil {
		return "", err
	}
	showId = uuid.String()

	return showId, nil
}
