package rCassandra

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	m "github.com/jovanfrandika/gochita/domain"
)

func (r *repository) GetMangaPostsByRange(ctx context.Context, start, end time.Time) (mangaPosts []m.DbMangaPost, err error) {
	mangaPosts = []m.DbMangaPost{}
	iter := r.session.Query(queryGetMangaPostsByRange, start, end).WithContext(ctx).Iter()
	var mangaPost m.DbMangaPost
	for iter.Scan(&mangaPost.Id, &mangaPost.Title, &mangaPost.Ref, &mangaPost.PublishedAt) {
		mangaPosts = append(mangaPosts, mangaPost)
	}
	if err = iter.Close(); err != nil {
		return []m.DbMangaPost{}, err
	}

	return mangaPosts, err
}

func (r *repository) GetMangaPostByTitle(ctx context.Context, title string) (mangaPost m.DbMangaPost, err error) {
	mangaPost = m.DbMangaPost{}
	err = r.session.Query(queryGetMangaPostByTitle, title).Consistency(gocql.One).WithContext(ctx).Scan(&mangaPost.Id, &mangaPost.Title, &mangaPost.Ref, &mangaPost.PublishedAt)
	return mangaPost, err
}

func (r *repository) CreateMangaPost(ctx context.Context, mangaPost m.FeedMangaPost) (mangaPostId string, err error) {
	uuid := gocql.TimeUUID()
	err = r.session.Query(queryCreateMangaPost, uuid, mangaPost.Title, mangaPost.Ref, mangaPost.PubDate).WithContext(ctx).Exec()
	if err != nil {
		return "", err
	}
	mangaPostId = uuid.String()

	return mangaPostId, nil
}
