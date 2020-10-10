package eventinputtypes

import "github.com/Jblew/cairparavel/functions/app/domain"

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
		NumberValue int64 `json:"integerValue"`
	} `json:"votingTime"`
	StartTime struct {
		NumberValue int64 `json:"integerValue"`
	} `json:"startTime"`
	EndTime struct {
		NumberValue int64 `json:"integerValue"`
	} `json:"endTime"`
	TimeConfirmed struct {
		BooleanValue bool `json:"booleanValue"`
	} `json:"timeConfirmed"`
	SignupTime struct {
		NumberValue int64 `json:"integerValue"`
	} `json:"signupTime"`
	MinParticipants struct {
		NumberValue int `json:"integerValue"`
	} `json:"minParticipants"`
	MaxParticipants struct {
		NumberValue int `json:"integerValue"`
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
				NumberValue int64 `json:"integerValue"`
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
		allowedTimes = append(allowedTimes, allowedTime.NumberValue)
	}

	return domain.Event{
		ID:               input.ID.StringValue,
		OwnerUID:         input.OwnerUID.StringValue,
		OwnerDisplayName: input.OwnerDisplayName.StringValue,
		VotingTime:       input.VotingTime.NumberValue,
		StartTime:        input.StartTime.NumberValue,
		EndTime:          input.EndTime.NumberValue,
		TimeConfirmed:    input.TimeConfirmed.BooleanValue,
		SignupTime:       input.SignupTime.NumberValue,
		MinParticipants:  input.MinParticipants.NumberValue,
		MaxParticipants:  input.MaxParticipants.NumberValue,
		Name:             input.Name.StringValue,
		Description:      input.Description.StringValue,
		AllowedTimes:     allowedTimes,
		CanSuggestTime:   input.CanSuggestTime.BooleanValue,
	}
}
