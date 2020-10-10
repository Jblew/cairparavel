package domain

import "time"

const (
	// EventStateTimeVoting — voting
	EventStateTimeVoting = iota
	// EventStateWaitingForTimeConfirm waiting for confirmation of time by event owner
	EventStateWaitingForTimeConfirm
	// EventStateCancelled cancelled
	EventStateCancelled
	// EventStateMembersSignup signing up
	EventStateMembersSignup
	// EventStateSignupClosed closed, waiting for the event to start
	EventStateSignupClosed
	// EventStateInProggress in proggress
	EventStateInProggress
	//EventStateFinished finished
	EventStateFinished
)

// EventState = state of the event
type EventState int

// GetEventStateAt retrives state of an event at any given time
func GetEventStateAt(event Event, atTime time.Time) EventState {
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
