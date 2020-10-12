package domain

import (
	"fmt"
	"time"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// Event an event organised at some time
type Event struct {
	ID               string  `json:"id"`
	OwnerUID         string  `json:"ownerUid" validate:"nonzero"`
	OwnerDisplayName string  `json:"ownerDisplayName" validate:"nonzero"`
	VotingTime       int64   `json:"votingTime"`
	StartTime        int64   `json:"startTime"`
	EndTime          int64   `json:"endTime"`
	TimeConfirmed    bool    `json:"timeConfirmed"`
	SignupTime       int64   `json:"signupTime"`
	MinParticipants  int     `json:"minParticipants"`
	MaxParticipants  int     `json:"maxParticipants"`
	Name             string  `json:"name" validate:"nonzero"`
	Description      string  `json:"description" validate:"nonzero"`
	AllowedTimes     []int64 `json:"allowedTimes"`
	CanSuggestTime   bool    `json:"canSuggestTime"`
}

// Validate validates
func (event Event) Validate(requireID bool) error {
	if requireID && len(event.ID) == 0 {
		return fmt.Errorf("ID is required on Event")
	}
	return validator.Validate(event)
}

// GetStateAt retrives state of an event at any given time
func (event *Event) GetStateAt(atTime time.Time, container *ioccontainer.Container) (EventState, error) {
	err := event.Validate(true)
	if err != nil {
		return EventStateNonexistent, err
	}

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
func (event *Event) OnStateChanged(previousState EventState, container *ioccontainer.Container) error {
	err := event.Validate(true)
	if err != nil {
		return err
	}

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
func (event *Event) OnCreated(container *ioccontainer.Container) error {
	err := event.Validate(true)
	if err != nil {
		return err
	}

	err = event.Observe(event.OwnerUID, container)
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
func (event *Event) OnModified(container *ioccontainer.Container) error {
	err := event.Validate(true)
	if err != nil {
		return err
	}

	payload := make(map[string]interface{})
	payload["event"] = event

	return event.NotifyObservers(Notification{
		Template: "event_modified",
		Payload:  payload,
	}, container)
}

// NotifyOwner notifies owner of the event
func (event *Event) NotifyOwner(notification Notification, container *ioccontainer.Container) error {
	var notificationQueue NotificationQueue
	container.Make(&notificationQueue)

	err := notificationQueue.Add(event.OwnerUID, notification)
	if err != nil {
		return err
	}
	return nil
}

// NotifyObservers notifies people observing the event
func (event *Event) NotifyObservers(notification Notification, container *ioccontainer.Container) error {
	var observersRepo EventObserverRepository
	container.Make(&observersRepo)
	var notificationQueue NotificationQueue
	container.Make(&notificationQueue)

	observers, err := observersRepo.GetAllForEvent(event.ID)
	if err != nil {
		return err
	}

	var lastErr error
	for _, observer := range observers {
		err := notificationQueue.Add(observer.UID, notification)
		if err != nil {
			lastErr = err
		}
	}
	return lastErr
}

// Observe observes an event
func (event *Event) Observe(userID string, container *ioccontainer.Container) error {
	var observersRepo EventObserverRepository
	container.Make(&observersRepo)

	return observersRepo.Add(EventObserver{
		EventID: event.ID,
		UID:     userID,
	})
}

// EventState = state of the event
type EventState string

const (
	// EventStateNonexistent — nonexistent, means event didn't exist before
	EventStateNonexistent EventState = "EventStateNonexistent"
	// EventStateTimeVoting — voting
	EventStateTimeVoting = "EventStateTimeVoting"
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
	GetAllNonFinishedAt(time time.Time) ([]Event, error)
}
