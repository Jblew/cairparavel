package domain

import "github.com/Jblew/ioccontainer/pkg/ioccontainer"

// EventSignup people signing for an event
type EventSignup struct {
	UID         string `json:"uid"`
	EventID     string `json:"eventId"`
	DisplayName string `json:"displayName"`
}

// OnAdded handles added signup
func (signup *EventSignup) OnAdded(container *ioccontainer.Container) error {
	return signup.sendNotificationAndObserve("member_signed_in", container)
}

// OnDeleted handles added signup
func (signup *EventSignup) OnDeleted(container *ioccontainer.Container) error {
	return signup.sendNotificationAndObserve("member_signed_out", container)
}

func (signup *EventSignup) sendNotificationAndObserve(templateName string, container *ioccontainer.Container) error {
	var eventRepository EventRepository
	container.Make(&eventRepository)

	event, err := eventRepository.GetEventByID(signup.EventID)
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
	notification.Payload["signup"] = signup

	return event.NotifyObservers(notification, container)
}

// EventSignupRepository manages signups
type EventSignupRepository interface {
	GetCount(eventID string) (int, error)
}
