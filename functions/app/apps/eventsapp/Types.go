package eventsapp

// EventObserversRepository manages people observing event
type EventObserversRepository interface {
	GetEventsObserversUids(eventID string) ([]string, error)
}
