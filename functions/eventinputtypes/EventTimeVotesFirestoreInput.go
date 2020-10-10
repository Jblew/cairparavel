package eventinputtypes

import (
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// EventTimeVotesFirestoreInput — firestore event input for EventTimeVotes
type EventTimeVotesFirestoreInput struct {
	UID struct {
		StringValue string `json:"stringValue"`
	} `json:"uid"`
	EventID struct {
		StringValue string `json:"stringValue"`
	} `json:"eventId"`
	DisplayName struct {
		StringValue string `json:"displayName"`
	} `json:"eventId"`
	Times struct {
		ArrayValue struct {
			Values []struct {
				NumberValue string `json:"integerValue"`
			} `json:"values"`
		} `json:"arrayValue"`
	} `json:"allowedTimes"`
}

// ToEventTimeVotes converter
func (input *EventTimeVotesFirestoreInput) ToEventTimeVotes() domain.EventTimeVotes {
	times := make([]int64, 0)
	for _, time := range input.Times.ArrayValue.Values {
		times = append(times, parseIntOrZero(time.NumberValue, 10, 64))
	}

	return domain.EventTimeVotes{
		UID:         input.UID.StringValue,
		EventID:     input.EventID.StringValue,
		DisplayName: input.DisplayName.StringValue,
		Times:       times,
	}
}
