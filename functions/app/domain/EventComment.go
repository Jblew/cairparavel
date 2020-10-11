package domain

import "github.com/golobby/container/pkg/container"

// EventComment is a comment on an event
type EventComment struct {
	ID        string `json:"id"`
	EventID   string `json:"eventId"`
	AuthorUID string `json:"authorUid"`
	Contents  string `json:"contents"`
	Time      int64  `json:"time"`
}

// OnAdded handles comment added to event
func (comment *EventComment) OnAdded(container container.Container) error {
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
