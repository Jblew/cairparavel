package eventinputtypes

import (
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// EventCommentFirestoreInput — firestore event input for EventComment
type EventCommentFirestoreInput struct {
	ID struct {
		StringValue string `json:"stringValue"`
	} `json:"id"`
	EventID struct {
		StringValue string `json:"stringValue"`
	} `json:"eventId"`
	AuthorUID struct {
		StringValue string `json:"stringValue"`
	} `json:"authorUid"`
	Contents struct {
		StringValue string `json:"stringValue"`
	} `json:"contents"`
	Time struct {
		NumberValue string `json:"integerValue"`
	} `json:"time"`
}

// ToEventComment converter
func (input *EventCommentFirestoreInput) ToEventComment() domain.EventComment {
	return domain.EventComment{
		ID:        input.ID.StringValue,
		EventID:   input.EventID.StringValue,
		AuthorUID: input.AuthorUID.StringValue,
		Contents:  input.Contents.StringValue,
		Time:      parseIntOrZero(input.Time.NumberValue, 10, 64),
	}
}
