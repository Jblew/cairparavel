package eventinputtypes

import (
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// EventFirestoreInput — firestore event input for Event
type EventFirestoreInput struct {
	ID struct {
		StringValue string `json:"stringValue"`
	} `json:"id"`
	OwnerUID struct {
		StringValue string `json:"stringValue"`
	} `json:"ownerUid"`
	OwnerDisplayName struct {
		StringValue string `json:"stringValue"`
	} `json:"ownerDisplayName"`
	VotingTime struct {
		NumberValue string `json:"integerValue"`
	} `json:"votingTime"`
	StartTime struct {
		NumberValue string `json:"integerValue"`
	} `json:"startTime"`
	EndTime struct {
		NumberValue string `json:"integerValue"`
	} `json:"endTime"`
	TimeConfirmed struct {
		BooleanValue bool `json:"booleanValue"`
	} `json:"timeConfirmed"`
	SignupTime struct {
		NumberValue string `json:"integerValue"`
	} `json:"signupTime"`
	MinParticipants struct {
		NumberValue string `json:"integerValue"`
	} `json:"minParticipants"`
	MaxParticipants struct {
		NumberValue string `json:"integerValue"`
	} `json:"maxParticipants"`
	Name struct {
		StringValue string `json:"stringValue"`
	} `json:"name"`
	Description struct {
		StringValue string `json:"stringValue"`
	} `json:"description"`
	AllowedTimes struct {
		ArrayValue struct {
			Values []struct {
				NumberValue string `json:"integerValue"`
			} `json:"values"`
		} `json:"arrayValue"`
	} `json:"allowedTimes"`
	CanSuggestTime struct {
		BooleanValue bool `json:"booleanValue"`
	} `json:"canSuggestTime"`
}

// ToEvent converter
func (input *EventFirestoreInput) ToEvent() domain.Event {
	allowedTimes := make([]int64, 0)

	for _, allowedTime := range input.AllowedTimes.ArrayValue.Values {
		allowedTimes = append(allowedTimes, parseIntOrZero(allowedTime.NumberValue, 10, 64))
	}

	return domain.Event{
		ID:               input.ID.StringValue,
		OwnerUID:         input.OwnerUID.StringValue,
		OwnerDisplayName: input.OwnerDisplayName.StringValue,
		VotingTime:       parseIntOrZero(input.VotingTime.NumberValue, 10, 64),
		StartTime:        parseIntOrZero(input.StartTime.NumberValue, 10, 64),
		EndTime:          parseIntOrZero(input.EndTime.NumberValue, 10, 64),
		TimeConfirmed:    input.TimeConfirmed.BooleanValue,
		SignupTime:       parseIntOrZero(input.SignupTime.NumberValue, 10, 64),
		MinParticipants:  int(parseIntOrZero(input.MinParticipants.NumberValue, 10, 32)),
		MaxParticipants:  int(parseIntOrZero(input.MaxParticipants.NumberValue, 10, 32)),
		Name:             input.Name.StringValue,
		Description:      input.Description.StringValue,
		AllowedTimes:     allowedTimes,
		CanSuggestTime:   input.CanSuggestTime.BooleanValue,
	}
}
