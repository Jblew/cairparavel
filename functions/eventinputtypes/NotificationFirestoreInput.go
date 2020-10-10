package eventinputtypes

import (
	"encoding/json"

	"github.com/Jblew/cairparavel/functions/app/domain"
)

// NotificationFirestoreInput — firestore event input for Notification
type NotificationFirestoreInput struct {
	ID struct {
		StringValue string `json:"stringValue"`
	} `json:"id"`
	UID struct {
		StringValue string `json:"stringValue"`
	} `json:"uid"`
	Template struct {
		StringValue string `json:"stringValue"`
	} `json:"template"`
	Payload struct {
		StringValue string `json:"stringValue"`
	} `json:"payload"`
}

// ToNotification converter
func (input *NotificationFirestoreInput) ToNotification() (domain.Notification, error) {
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(input.Payload.StringValue), &payload)
	if err != nil {
		return domain.Notification{}, err
	}

	return domain.Notification{
		ID:       input.ID.StringValue,
		UID:      input.UID.StringValue,
		Template: input.Template.StringValue,
		Payload:  payload,
	}, nil
}
