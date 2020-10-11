package domain

import (
	"time"

	"log"

	"github.com/golobby/container/pkg/container"
)

// Event an event organised at some time
type Event struct {
	ID               string  `json:"id"`
	OwnerUID         string  `json:"ownerUid"`
	OwnerDisplayName string  `json:"ownerDisplayName"`
	VotingTime       int64   `json:"votingTime"`
	StartTime        int64   `json:"startTime"`
	EndTime          int64   `json:"endTime"`
	TimeConfirmed    bool    `json:"timeConfirmed"`
	SignupTime       int64   `json:"signupTime"`
	MinParticipants  int     `json:"minParticipants"`
	MaxParticipants  int     `json:"maxParticipants"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	AllowedTimes     []int64 `json:"allowedTimes"`
	CanSuggestTime   bool    `json:"canSuggestTime"`
}

// GetStateAt retrives state of an event at any given time
func (event *Event) GetStateAt(atTime time.Time, container container.Container) (EventState, error) {
	var signupRepo EventSignupRepository
	container.Make(&signupRepo)

	timeMs := atTime.UnixNano() / int64(time.Millisecond)
	if timeMs < event.VotingTime {
		return EventStateTimeVoting, nil
	} else if event.TimeConfirmed == false {
		if timeMs < event.StartTime {
			return EventStateWaitingForTimeConfirm, nil
		}
		return EventStateCancelled, nil
	} else if timeMs < event.SignupTime {
		return EventStateMembersSignup, nil
	}

	signupCount, err := signupRepo.GetCount(event.ID)
	if err != nil {
		return EventStateCancelled, err
	}
	if signupCount < event.MinParticipants {
		return EventStateCancelled, nil
	} else if timeMs < event.StartTime {
		return EventStateSignupClosed, nil
	} else if timeMs < event.EndTime {
		return EventStateInProggress, nil
	}
	return EventStateFinished, nil
}

// OnStateChanged handler
func (event *Event) OnStateChanged(previousState EventState, container container.Container) error {
	payload := make(map[string]interface{})
	payload["event"] = event
	payload["previousState"] = previousState

	eventState, err := event.GetStateAt(time.Now(), container)
	if err != nil {
		return err
	}

	if eventState == EventStateTimeVoting {
		return event.NotifyObservers(Notification{
			Template: "event_voting_started",
			Payload:  payload,
		}, container)
	} else if eventState == EventStateMembersSignup {
		return event.NotifyObservers(Notification{
			Template: "event_members_signup_started",
			Payload:  payload,
		}, container)
	} else if eventState == EventStateSignupClosed {
		return event.NotifyObservers(Notification{
			Template: "event_members_signup_closed",
			Payload:  payload,
		}, container)
	} else if eventState == EventStateInProggress {
		return event.NotifyObservers(Notification{
			Template: "event_started",
			Payload:  payload,
		}, container)
	} else if eventState == EventStateFinished {
		return event.NotifyObservers(Notification{
			Template: "event_cancelled",
			Payload:  payload,
		}, container)
	}
	return nil
}

// OnCreated handler
func (event *Event) OnCreated(container container.Container) error {
	err := event.Observe(event.OwnerUID, container)
	if err != nil {
		return err
	}

	payload := make(map[string]interface{})
	payload["event"] = event

	return event.NotifyObservers(Notification{
		Template: "event_created",
		Payload:  payload,
	}, container)
}

// OnModified handler
func (event *Event) OnModified(container container.Container) error {
	payload := make(map[string]interface{})
	payload["event"] = event

	return event.NotifyObservers(Notification{
		Template: "event_modified",
		Payload:  payload,
	}, container)
}

// NotifyOwner notifies owner of the event
func (event *Event) NotifyOwner(notification Notification, container container.Container) error {
	var notificationQueue NotificationQueue
	container.Make(&notificationQueue)

	err := notificationQueue.Add(event.OwnerUID, notification)
	if err != nil {
		return err
	}
	return nil
}

// NotifyObservers notifies people observing the event
func (event *Event) NotifyObservers(notification Notification, container container.Container) error {
	var observersRepo EventObserverRepository
	container.Make(&observersRepo)

	var notificationQueue NotificationQueue
	container.Make(&notificationQueue)

	observers, err := observersRepo.GetEventObservers(event.ID)
	if err != nil {
		return err
	}

	var lastErr error
	for _, observer := range observers {
		err := notificationQueue.Add(observer.UID, notification)
		if err != nil {
			log.Printf("Error while sending notification %v", err)
			lastErr = err
		}
	}
	return lastErr
}

// Observe observes an event
func (event *Event) Observe(userID string, container container.Container) error {
	var observersRepo EventObserverRepository
	container.Make(&observersRepo)

	return observersRepo.AddEventObserver(EventObserver{
		EventID: event.ID,
		UID:     userID,
	})
}

// EventState = state of the event
type EventState string

const (
	// EventStateTimeVoting — voting
	EventStateTimeVoting EventState = "EventStateTimeVoting"
	// EventStateWaitingForTimeConfirm waiting for confirmation of time by event owner
	EventStateWaitingForTimeConfirm = "EventStateWaitingForTimeConfirm"
	// EventStateCancelled cancelled
	EventStateCancelled = "EventStateCancelled"
	// EventStateMembersSignup signing up
	EventStateMembersSignup = "EventStateMembersSignup"
	// EventStateSignupClosed closed, waiting for the event to start
	EventStateSignupClosed = "EventStateSignupClosed"
	// EventStateInProggress in proggress
	EventStateInProggress = "EventStateInProggress"
	//EventStateFinished finished
	EventStateFinished = "EventStateFinished"
)

// EventRepository is a repository for events
type EventRepository interface {
	GetEventByID(ID string) (Event, error)
}
