package faq

type messages struct {
	Failed      string
	InvalidUUID string
	NotFound    string
}

var i18nMessages = messages{
	Failed:      "feedback_failed",
	InvalidUUID: "feedback_invalid_uuid",
	NotFound:    "feedback_not_found",
}
