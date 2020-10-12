package domain

import "gopkg.in/validator.v2"

// EventObserver is user observing an event
type EventObserver struct {
	EventID string `json:"eventId" validate:"nonzero"`
	UID     string `json:"uid" validate:"nonzero"`
}

// Validate validates
func (observer EventObserver) Validate() error {
	return validator.Validate(observer)
}

// EventObserverRepository is a repository of event observers
type EventObserverRepository interface {
	GetAllForEvent(eventID string) ([]EventObserver, error)
	Add(observer EventObserver) error
}
