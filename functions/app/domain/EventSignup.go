package domain

// EventSignup people signing for an event
type EventSignup struct {
	UID         string `json:"uid"`
	EventID     string `json:"eventId"`
	DisplayName string `json:"displayName"`
}
