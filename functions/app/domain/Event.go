package domain

import "time"

// Event an event organised at some time
type Event struct {
	ID               string                    `json:"id"`
	OwnerUID         string                    `json:"ownerUid"`
	OwnerDisplayName string                    `json:"ownerDisplayName"`
	VotingTime       int64                     `json:"votingTime"`
	StartTime        int64                     `json:"startTime"`
	EndTime          int64                     `json:"endTime"`
	TimeConfirmed    bool                      `json:"timeConfirmed"`
	SignupTime       int64                     `json:"signupTime"`
	Votes            map[string]EventTimeVotes `json:"votes"`
	SignedMembers    map[string]EventSignup    `json:"signedMembers"`
	MinParticipants  int                       `json:"minParticipants"`
	MaxParticipants  int                       `json:"maxParticipants"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description"`
	AllowedTimes     []int64                   `json:"allowedTimes"`
	CanSuggestTime   bool                      `json:"canSuggestTime"`
}

// EventTimeVotes votes for organising event at some specific time
type EventTimeVotes struct {
	UID         string  `json:"uid"`
	DisplayName string  `json:"displayName"`
	Times       []int64 `json:"times"`
}

// EventSignup people signing for an event
type EventSignup struct {
	UID         string `json:"uid"`
	DisplayName string `json:"displayName"`
}

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
