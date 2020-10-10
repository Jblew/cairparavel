package domain

// EventTimeVotes votes for organising event at some specific time
type EventTimeVotes struct {
	UID         string  `json:"uid"`
	EventID     string  `json:"eventId"`
	DisplayName string  `json:"displayName"`
	Times       []int64 `json:"times"`
}
