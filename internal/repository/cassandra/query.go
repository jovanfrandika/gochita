package rCassandra

const (
	queryGetShowByTitle = `
		SELECT id, title, thumbnail, category, ref
		FROM show
		WHERE title = ?
		LIMIT 1
	`
	queryCreateShow = `
		INSERT INTO show (id, title, thumbnail, category, ref)
		VALUES (?, ?, ?, ?, ?)
	`
	queryGetShowEpisodesByShowId = `
		SELECT id, showId, num, publishedAt
		FROM showEpisode
		WHERE showId = ?
	`
	queryCreateShowEpisode = `
		INSERT INTO showEpisode (id, showId, num, publishedAt)
		VALUES (?, ?, ?, ?)
	`
)
