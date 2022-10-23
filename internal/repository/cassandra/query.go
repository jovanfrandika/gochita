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
		SELECT id, show_id, num, published_at
		FROM show_episode
		WHERE show_id = ?
	`
	queryGetShowEpisodesByRange = `
		SELECT id, show_id, num, published_at
		FROM show_episode
		WHERE published_at >= ? AND published_at <= ? 
		LIMIT 100
		ALLOW FILTERING
	`
	queryCreateShowEpisode = `
		INSERT INTO show_episode (id, show_id, num, published_at)
		VALUES (?, ?, ?, ?)
	`

	// Channel Subscription
	queryGetSubscriptionsByReferenceId = `
		SELECT id, subscription_type, reference_id, context_id, is_enabled
		FROM channel_subscription
		WHERE subscription_type = ? AND reference_id = ? AND is_enabled = ?
		ALLOW FILTERING
	`
	queryGetSubscriptionsByContextId = `
		SELECT id, subscription_type, reference_id, context_id, is_enabled
		FROM channel_subscription
		WHERE subscription_type = ? AND context_id = ?  AND is_enabled = ?
		ALLOW FILTERING
	`
	queryGetSubscriptions = `
		SELECT id, subscription_type, reference_id, context_id, is_enabled
		FROM channel_subscription
		WHERE subscription_type = ? AND is_enabled = ?
		LIMIT 1000
		ALLOW FILTERING
	`
	queryGetSubscription = `
		SELECT id, subscription_type, reference_id, context_id, is_enabled
		FROM channel_subscription
		WHERE subscription_type = ? AND reference_id = ? AND context_id  = ?
		LIMIT 1
	`
	queryCreateSubscription = `
		INSERT INTO channel_subscription (id, subscription_type, reference_id, context_id, is_enabled)
		VALUES (?, ?, ?, ?, True)
	`
	queryToggleSubscriptions = `
		UPDATE channel_subscription
		SET is_enabled = ?
		WHERE subscription_type = ? AND reference_id = ? AND context_id IN ?
	`
	queryToggleSubscription = `
		UPDATE channel_subscription
		SET is_enabled = ?
		WHERE subscription_type = ? AND reference_id = ? AND context_id = ?
		IF EXISTS
	`

	// ChannelNotificiation
	queryGetNotification = `
		SELECT channel_subscription_id, notified_at
		FROM channel_notification
		WHERE channel_subscription_id = ?
		LIMIT 1
	`
	queryCreateNotification = `
		INSERT INTO channel_notification (channel_subscription_id, notified_at)
		VALUES (?, ?)
	`

	// Headline
	queryGetHeadlinesByRange = `
		SELECT id, title, thumbnail, ref, published_at
		FROM headline
		WHERE published_at >= ? AND published_at <= ? 
		LIMIT 100
		ALLOW FILTERING
	`
	queryGetHeadlineByTitle = `
		SELECT id, title, thumbnail, ref, published_at
		FROM headline
		WHERE title = ?
		LIMIT 1
	`
	queryCreateHeadline = `
		INSERT INTO headline (id, title, thumbnail, ref, published_at)
		VALUES (?, ?, ?, ?, ?)
	`

	// MangaPost
	queryGetMangaPostsByRange = `
		SELECT id, title, ref, published_at
		FROM manga_post
		WHERE published_at >= ? AND published_at <= ? 
		LIMIT 100
		ALLOW FILTERING
	`
	queryGetMangaPostByTitle = `
		SELECT id, title, ref, published_at
		FROM manga_post
		WHERE title = ?
		LIMIT 1
	`
	queryCreateMangaPost = `
		INSERT INTO manga_post (id, title, ref, published_at)
		VALUES (?, ?, ?, ?)
	`
)
