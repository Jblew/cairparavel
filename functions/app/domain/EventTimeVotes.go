package domain

import "github.com/golobby/container/pkg/container"

// EventTimeVotes votes for organising event at some specific time
type EventTimeVotes struct {
	UID         string  `json:"uid"`
	EventID     string  `json:"eventId"`
	DisplayName string  `json:"displayName"`
	Times       []int64 `json:"times"`
}

// OnAdded handles added vote
func (votes *EventTimeVotes) OnAdded(container container.Container) error {
	return votes.sendNotification("event_voted", container)
}

// OnModified handles added vote
func (votes *EventTimeVotes) OnModified(container container.Container) error {
	return votes.sendNotification("event_vote_modified", container)
}

// OnDeleted handles added vote
func (votes *EventTimeVotes) OnDeleted(container container.Container) error {
	return votes.sendNotification("event_vote_deleted", container)
}

func (votes *EventTimeVotes) sendNotification(templateName string, container container.Container) error {
	var eventRepository EventRepository
	container.Make(&eventRepository)

	event, err := eventRepository.GetEventByID(votes.EventID)
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
