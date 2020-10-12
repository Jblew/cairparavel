package domain

import (
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// EventTimeVotes votes for organising event at some specific time
type EventTimeVotes struct {
	UID         string  `json:"uid" validate:"nonzero"`
	EventID     string  `json:"eventId" validate:"nonzero"`
	DisplayName string  `json:"displayName" validate:"nonzero"`
	Times       []int64 `json:"times"`
}

// Validate validates
func (votes EventTimeVotes) Validate() error {
	return validator.Validate(votes)
}

// OnAdded handles added vote
func (votes *EventTimeVotes) OnAdded(container *ioccontainer.Container) error {
	return votes.sendNotificationAndObserve("event_voted", container)
}

// OnModified handles added vote
func (votes *EventTimeVotes) OnModified(container *ioccontainer.Container) error {
	return votes.sendNotificationAndObserve("event_vote_modified", container)
}

// OnDeleted handles added vote
func (votes *EventTimeVotes) OnDeleted(container *ioccontainer.Container) error {
	return votes.sendNotificationAndObserve("event_vote_deleted", container)
}

func (votes *EventTimeVotes) sendNotificationAndObserve(templateName string, container *ioccontainer.Container) error {
	var eventRepository EventRepository
	container.Make(&eventRepository)

	event, err := eventRepository.GetEventByID(votes.EventID)
	if err != nil {
		return err
	}

	err = event.Observe(event.OwnerUID, container)
	if err != nil {
		return err
	}

	notification := Notification{
		Template: templateName,
		Payload:  make(map[string]interface{}),
	}

	notification.Payload["event"] = event
	notification.Payload["votes"] = votes

	return event.NotifyObservers(notification, container)
}
