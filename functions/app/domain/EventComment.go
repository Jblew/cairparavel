package domain

import (
	"fmt"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// EventComment is a comment on an event
type EventComment struct {
	ID        string `json:"id"`
	EventID   string `json:"eventId" validate:"nonzero"`
	AuthorUID string `json:"authorUid" validate:"nonzero"`
	Contents  string `json:"contents" validate:"nonzero"`
	Time      int64  `json:"time"`
}

// Validate validates
func (comment EventComment) Validate(requireID bool) error {
	if requireID && len(comment.ID) == 0 {
		return fmt.Errorf("ID is required on EventComment")
	}
	return validator.Validate(comment)
}

// OnAdded handles comment added to event
func (comment *EventComment) OnAdded(container *ioccontainer.Container) error {
	err := comment.Validate(true)
	if err != nil {
		return err
	}

	var eventRepository EventRepository
	container.Make(&eventRepository)

	event, err := eventRepository.GetEventByID(comment.EventID)
	if err != nil {
		return err
	}

	err = event.Observe(event.OwnerUID, container)
	if err != nil {
		return err
	}

	notification := Notification{
		Template: "comment_added",
		Payload:  make(map[string]interface{}),
	}

	notification.Payload["event"] = event
	notification.Payload["comment"] = comment

	return event.NotifyObservers(notification, container)
}
