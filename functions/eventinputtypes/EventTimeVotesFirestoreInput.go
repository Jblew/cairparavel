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
		StringValue string `json:"stringValue"`
	} `json:"displayName"`
	Times struct {
		ArrayValue struct {
			Values []struct {
				NumberValue string `json:"integerValue"`
			} `json:"values"`
		} `json:"arrayValue"`
	} `json:"times"`
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

/*
map[
	displayName:map[stringValue:Jedrzej Lew]
	eventId:map[stringValue:pYqj0CxB23S1bq0yhVr7]
	times:map[
		arrayValue:map[values:[map[integerValue:1603402133890] map[integerValue:1603315733890]]
	]] uid:map[stringValue:WuHBEm64OxQMqc5PHqN9rVj1Inm2]]
*/
