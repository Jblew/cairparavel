package eventsapp

import "github.com/Jblew/cairparavel/functions/app/domain"

// EventObserversRepository manages people observing event
type EventObserversRepository interface {
	GetEventsObserversUids(eventID string) ([]string, error)
}

// EventStateChangedRepository is a repository for getting and updating event state to check if it changed
type EventStateChangedRepository interface {
	GetPreviousAndUpdateEventState(eventID string, state domain.EventState) (domain.EventState, error)
}
