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
	return signup.sendNotification("event_voted", container)
}

// OnDeleted handles added vote
func (votes *EventTimeVotes) OnDeleted(container container.Container) error {
	return signup.sendNotification("event_vote_deleted", container)
}

func (signup *EventSignup) sendNotification(templateName string, container container.Container) error {
	var eventRepository EventRepository
	container.Make(&eventRepository)

	event, err := eventRepository.GetEventByID(signup.EventID)
	if err != nil {
		return err
	}

	notification := Notification{
		Template: templateName,
		Payload:  make(map[string]interface{}),
	}

	payload["event"] = event
	payload["votes"] = votes

	return event.NotifyObservers(notification, container)
}
