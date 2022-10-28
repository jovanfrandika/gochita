package rCassandra

import (
	"context"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/gochita/domain"
)

func (r *repository) GetShowById(ctx context.Context, showId string) (show m.DbShow, err error) {
	show = m.DbShow{}
	err = r.session.Query(queryGetShowById, showId).Consistency(gocql.One).WithContext(ctx).Scan(&show.Id, &show.Title, &show.Thumbnail, &show.Category, &show.Ref)
	return show, err
}

func (r *repository) GetShowByTitle(ctx context.Context, title string) (show m.DbShow, err error) {
	show = m.DbShow{}
	err = r.session.Query(queryGetShowByTitle, title).Consistency(gocql.One).WithContext(ctx).Scan(&show.Id, &show.Title, &show.Thumbnail, &show.Category, &show.Ref)
	return show, err
}

func (r *repository) CreateShow(ctx context.Context, show m.FeedShow) (showId string, err error) {
	uuid := gocql.TimeUUID()
	err = r.session.Query(queryCreateShow, uuid, show.Title, show.Thumbnail, show.Category, show.Ref).WithContext(ctx).Exec()
	if err != nil {
		return "", err
	}
	showId = uuid.String()

	return showId, nil
}
