package domain

// EventObserver is user observing an event
type EventObserver struct {
	EventID string `json:"eventId"`
	UID     string `json:"uid"`
}

// EventObserverRepository is a repository of event observers
type EventObserverRepository interface {
	GetEventObservers(eventID string) ([]EventObserver, error)
}
