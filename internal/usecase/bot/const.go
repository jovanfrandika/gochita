package uBot

const (
	LABEL_NEW_SHOW_MOVIE_EPISODE  = "***Reminder!!!***\nTitle: %v\nPublished Date: %v\n"
	LABEL_NEW_SHOW_SERIES_EPISODE = "***Reminder!!!***\nTitle: %v\nEpisode: %v\nPublished Date: %v\n"
	LABEL_NEW_HEADLINE            = "***News!!!***\nTitle: %v\nPublished Date: %v\nSauce: %v\n"

	LABEL_SUBSCRIPTION_TITLE = "%v. %v\n"

	LABEL_SUCCESS_NEW_SHOW_SUBSCRIPTION   = "Successfully subscribed to new shows"
	LABEL_UNSUCCESS_NEW_SHOW_SUBSCRIPTION = "Subscription failed :("

	LABEL_SUCCESS_NEW_SHOW_UNSUBSCRIPTION   = "Successfully unsubscribed to new shows"
	LABEL_UNSUCCESS_NEW_SHOW_UNSUBSCRIPTION = "Unsubscription failed :("

	LABEL_SUCCESS_SPECIFIC_SHOW_SUBSCRIPTION   = "%v successfully subscribed!"
	LABEL_UNSUCCESS_SPECIFIC_SHOW_SUBSCRIPTION = "%v subscription failed :("

	LABEL_SUCCESS_SPECIFIC_SHOW_UNSUBSCRIPTION   = "%v successfully unsubscribed!"
	LABEL_UNSUCCESS_SPECIFIC_SHOW_UNSUBSCRIPTION = "%v unsubscription failed :("

	LABEL_SUCCESS_NEW_HEADLINE_SUBSCRIPTION   = "Headlines successfully subscribed!"
	LABEL_UNSUCCESS_NEW_HEADLINE_SUBSCRIPTION = "Headlines subscription failed :("

	LABEL_SUCCESS_NEW_HEADLINE_UNSUBSCRIPTION   = "Headlines successfully unsubscribed!"
	LABEL_UNSUCCESS_NEW_HEADLINE_UNSUBSCRIPTION = "Headlines unsubscription failed :("

	NO_SUBSCRIPTIONS = "No subscriptions"
	DEFAULT_ERROR    = "Oops something went wrong!"

	SUBSCRIPTION_TYPE_NEW_SHOW          = 1
	SUBSCRIPTION_TYPE_SPECIFIC_SHOW     = 2
	SUBSCRIPTION_TYPE_NEW_HEADLINE      = 3
	SUBSCRIPTION_TYPE_SPECIFIC_HEADLINE = 4

	NO_CONTEXT_ID = "00000000-0000-0000-0000-000000000000"
)
