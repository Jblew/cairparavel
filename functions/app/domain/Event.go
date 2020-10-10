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
func (event *Event) GetStateAt(atTime time.Time) EventState {
	timeMs := atTime.UnixNano() / int64(time.Millisecond)
	if timeMs < event.VotingTime {
		return EventStateTimeVoting
	} else if event.TimeConfirmed == false {
		if timeMs < event.StartTime {
			return EventStateWaitingForTimeConfirm
		}
		return EventStateCancelled
	} else if timeMs < event.SignupTime {
		return EventStateMembersSignup
	} else if len(event.SignedMembers) < event.MinParticipants {
		return EventStateCancelled
	} else if timeMs < event.StartTime {
		return EventStateSignupClosed
	} else if timeMs < event.EndTime {
		return EventStateInProggress
	}
	return EventStateFinished
}

// OnStateChanged handler
func (event *Event) OnStateChanged(previousState EventState, container container.Container) error {
	payload := make(map[string]interface{})
	payload["event"] = event
	payload["previousState"] = previousState

	eventState := GetEventStateAt(event, time.Now())
	if eventState == EventStateTimeVoting {
		return event.NotifyObservers(Notification{
			Template: "event_voting_started",
			Payload:  payload,
		})
	} else if eventState == EventStateMembersSignup {
		return event.NotifyObservers(Notification{
			Template: "event_members_signup_started",
			Payload:  payload,
		})
	} else if eventState == EventStateSignupClosed {
		return event.NotifyObservers(Notification{
			Template: "event_members_signup_closed",
			Payload:  payload,
		})
	} else if eventState == EventStateInProggress {
		return event.NotifyObservers(Notification{
			Template: "event_started",
			Payload:  payload,
		})
	} else if eventState == EventStateFinished {
		return event.NotifyObservers(Notification{
			Template: "event_cancelled",
			Payload:  payload,
		})
	}
	return nil
}

// NotifyObservers notifies people observing the event
func (event *Event) NotifyObservers(notification Notification, container container.Container) {
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
		err := notificationQueue.ScheduleToSend(observer.UID, notification)
		if err != nil {
			log.Printf("Error while sending notification %v", err)
			lastErr = err
		}
	}
	return lastErr
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
