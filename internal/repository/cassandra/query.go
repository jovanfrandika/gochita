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

	// ChannelShowSubscription
	queryGetShowSubscriptionsByReferenceId = `
		SELECT reference_id, show_id, is_enabled
		FROM channel_show_subscription
		WHERE reference_id = ? AND is_enabled = ?
		ALLOW FILTERING
	`
	queryGetShowSubscriptionsByShowId = `
		SELECT reference_id, show_id, is_enabled
		FROM channel_show_subscription
		WHERE show_id = ? AND is_enabled = ?
		ALLOW FILTERING
	`
	queryGetShowSubscription = `
		SELECT reference_id, show_id, is_enabled
		FROM channel_show_subscription
		WHERE reference_id = ? AND show_id = ?
		LIMIT 1
	`
	queryToggleShowSubscription = `
		UPDATE channel_show_subscription
		SET is_enabled = ?
		WHERE reference_id = ? AND show_id = ?
		IF EXISTS
	`

	// ChannelShowEpisodeNotificiation
	queryGetShowNotification = `
		SELECT reference_id, show_episode_id, notified_at
		FROM channel_show_episode_notification
		WHERE reference_id = ? AND show_episode_id = ?
		LIMIT 1
	`
	queryCreateShowNotification = `
		INSERT INTO channel_show_episode_notification (reference_id, show_episode_id, notified_at)
		VALUES (?, ?, ?)
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

	// Channel Headline Subscription
	queryGetHeadlineSubscriptions = `
		SELECT reference_id, is_enabled
		FROM channel_headline_subscription
		WHERE is_enabled = true
		LIMIT 1000
		ALLOW FILTERING
	`
	queryGetHeadlineSubscription = `
		SELECT reference_id, is_enabled
		FROM channel_headline_subscription
		WHERE reference_id = ?
		LIMIT 1
	`
	queryToggleHeadlineSubscription = `
		UPDATE channel_headline_subscription
		SET is_enabled = ?
		WHERE reference_id = ?
	`

	// Channel Headline Subscription
	queryGetHeadlineNotification = `
		SELECT reference_id, headline_id, notified_at
		FROM channel_headline_notification
		WHERE reference_id = ? AND headline_id = ?
		LIMIT 1
	`
	queryCreateHeadlineNotification = `
		INSERT INTO channel_headline_notification (reference_id, headline_id, notified_at)
		VALUES (?, ?, ?)
	`
)
