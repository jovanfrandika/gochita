package rCassandra

const (
	// Show
	queryGetShowById = `
		SELECT id, title, thumbnail, category, ref
		FROM show
		WHERE id = ?
		LIMIT 1
	`
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

	// ShowEpisode
	queryGetShowEpisodesByShowId = `
		SELECT id, showId, num, publishedAt
		FROM showEpisode
		WHERE showId = ?
	`
	queryGetShowEpisodesByRange = `
		SELECT id, showId, num, publishedAt
		FROM showEpisode
		WHERE publishedAt >= ? AND publishedAt <= ? 
		LIMIT 100
	`
	queryCreateShowEpisode = `
		INSERT INTO showEpisode (id, showId, num, publishedAt)
		VALUES (?, ?, ?, ?)
	`

	// ChannelShowSubscription
	queryGetSubscriptionsByReferenceId = `
		SELECT showId, referenceId, isEnabled
		FROM channelShowSubscription
		WHERE referenceId = ?
	`
	queryGetSubscriptionsByShowId = `
		SELECT showId, referenceId, isEnabled
		FROM channelShowSubscription
		WHERE showId = ?
	`
	queryGetSubscription = `
		SELECT showId, referenceId, isEnabled
		FROM channelShowSubscription
		WHERE referenceId = ? AND showId = ?
		LIMIT 1
	`
	queryCreateSubscription = `
		INSERT INTO channelShowSubscription (showId, referenceId, isEnabled)
		VALUES (?, ?, true)
	`
	queryToggleSubscription = `
		UPDATE channelShowSubscription
		SET isEnabled = ?
		WHERE referenceId = ? AND showId = ?
		IF EXISTS
	`

	// ChannelShowEpisodeNotificiation
	queryGetNotification = `
		SELECT showEpisodeId, referenceId, notifiedAt
		FROM channelShowEpisodeNotificiation
		WHERE referenceId = ? AND showEpisodeId = ?
		LIMIT 1
	`
	queryCreateNotification = `
		INSERT INTO channelShowEpisodeNotificiation (showEpisode, referenceId, notifiedAt)
		VALUES (?, ?, ?)
	`
)
