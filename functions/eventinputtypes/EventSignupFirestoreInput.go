package eventinputtypes

import "github.com/Jblew/cairparavel/functions/app/domain"

// EventSignupFirestoreInput — firestore event input for EventSignup
type EventSignupFirestoreInput struct {
	UID struct {
		StringValue string `json:"stringValue"`
	} `json:"uid"`
	EventID struct {
		StringValue string `json:"stringValue"`
	} `json:"eventId"`
	DisplayName struct {
		StringValue string `json:"stringValue"`
	} `json:"displayName"`
}

// ToEventSignup converter
func (input *EventSignupFirestoreInput) ToEventSignup() domain.EventSignup {
	return domain.EventSignup{
		UID:         input.UID.StringValue,
		EventID:     input.EventID.StringValue,
		DisplayName: input.DisplayName.StringValue,
	}
}
