package domain

type Notification struct {
	ID       string                 `json:"id"`
	UID      string                 `json:"uid"`
	Template string                 `json:"template"`
	Payload  map[string]interface{} `json:"payload"`
}

type EventObserversNotifier interface {
	NotifyEventObservers(event Event, notification Notification) error
}
