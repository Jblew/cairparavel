package domain

// EventComment is a comment on an event
type EventComment struct {
	ID        string `json:"id"`
	EventID   string `json:"eventId"`
	AuthorUID string `json:"authorUid"`
	Contents  string `json:"contents"`
	Time      int64  `json:"time"`
}
