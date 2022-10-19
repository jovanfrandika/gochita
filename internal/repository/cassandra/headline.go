package rCassandra

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func (r *repository) GetHeadlinesByRange(ctx context.Context, start, end time.Time) (headlines []m.DbHeadline, err error) {
	headlines = []m.DbHeadline{}
	iter := r.session.Query(queryGetHeadlinesByRange, start, end).Iter()
	var headline m.DbHeadline
	for iter.Scan(&headline.Id, &headline.Title, &headline.Thumbnail, &headline.Ref, &headline.PublishedAt) {
		headlines = append(headlines, headline)
	}
	if err = iter.Close(); err != nil {
		return []m.DbHeadline{}, err
	}

	return headlines, err
}

func (r *repository) GetHeadlineByTitle(ctx context.Context, title string) (headline m.DbHeadline, err error) {
	headline = m.DbHeadline{}
	err = r.session.Query(queryGetHeadlineByTitle, title).Consistency(gocql.One).Scan(&headline.Id, &headline.Title, &headline.Thumbnail, &headline.Ref, &headline.PublishedAt)
	return headline, err
}

func (r *repository) CreateHeadline(ctx context.Context, headline m.FeedHeadline) (headlineId string, err error) {
	uuid := gocql.TimeUUID()
	err = r.session.Query(queryCreateHeadline, uuid, headline.Title, headline.Thumbnail, headline.Ref, headline.PubDate).Exec()
	if err != nil {
		return "", err
	}
	headlineId = uuid.String()

	return headlineId, nil
}
