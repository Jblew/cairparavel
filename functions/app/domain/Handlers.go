package domain

import (
	"time"

	"github.com/golobby/container/pkg/container"
)

// OnEventStateChanged handler
func OnEventStateChanged(event Event, container container.Container) error {
	var eventObserversNotifier EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event

	eventState := GetEventStateAt(event, time.Now())
	if eventState == EventStateTimeVoting {
		return eventObserversNotifier.NotifyEventObservers(event, Notification{
			Template: "event_voting_started",
			Payload:  payload,
		})
	} else if eventState == EventStateMembersSignup {
		return eventObserversNotifier.NotifyEventObservers(event, Notification{
			Template: "event_members_signup_started",
			Payload:  payload,
		})
	} else if eventState == EventStateSignupClosed {
		return eventObserversNotifier.NotifyEventObservers(event, Notification{
			Template: "event_members_signup_closed",
			Payload:  payload,
		})
	} else if eventState == EventStateInProggress {
		return eventObserversNotifier.NotifyEventObservers(event, Notification{
			Template: "event_started",
			Payload:  payload,
		})
	} else if eventState == EventStateFinished {
		return eventObserversNotifier.NotifyEventObservers(event, Notification{
			Template: "event_cancelled",
			Payload:  payload,
		})
	}
	return nil
}

// OnMessengerMessage handler
func OnMessengerMessage(messageText string, recipient MessengerRecipient, container container.Container) error {
	var messengerNotifier MessengerNotifier
	container.Make(&messengerNotifier)

	payload := make(map[string]interface{})
	payload["messageText"] = messageText

	return messengerNotifier.SendNotification(recipient, Notification{
		Template: "messenger_respond",
		Payload:  payload,
	})
}

// OnMessengerReferral handler
func OnMessengerReferral(referralCode string, messengerRecipient MessengerRecipient, container container.Container) error {
	var messengerRecipientRepository MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)
	var messengerNotifier MessengerNotifier
	container.Make(&messengerNotifier)
	var usersRepository UsersRepository
	container.Make(&usersRepository)

	userID := referralCode
	err := messengerRecipientRepository.StoreMessengerRecipient(userID, messengerRecipient)
	if err != nil {
		return err
	}
	user, err := usersRepository.GetUser(userID)
	if err != nil {
		return err
	}

	payload := make(map[string]interface{})
	payload["user"] = user
	return messengerNotifier.SendNotification(messengerRecipient, Notification{
		Template: "messenger_welcome",
		Payload:  payload,
	})
}

// OnEventMemberSignup handler
func OnEventMemberSignup(event Event, signup EventSignup, container container.Container) error {
	var eventObserversNotifier EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["signup"] = signup
	return eventObserversNotifier.NotifyEventObservers(event, Notification{
		Template: "member_signed_in",
		Payload:  payload,
	})
}

// OnEventMemberSignout handler
func OnEventMemberSignout(event Event, signup EventSignup, container container.Container) error {
	var eventObserversNotifier EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["signup"] = signup
	return eventObserversNotifier.NotifyEventObservers(event, Notification{
		Template: "member_signed_out",
		Payload:  payload,
	})
}

// OnEventVote handler
func OnEventVote(event Event, votes EventTimeVotes, container container.Container) error {
	var eventObserversNotifier EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["votes"] = votes
	return eventObserversNotifier.NotifyEventObservers(event, Notification{
		Template: "event_voted",
		Payload:  payload,
	})
}

// OnEventVoteDeleted handler
func OnEventVoteDeleted(event Event, votes EventTimeVotes, container container.Container) error {
	var eventObserversNotifier EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["votes"] = votes
	return eventObserversNotifier.NotifyEventObservers(event, Notification{
		Template: "event_vote_deleted",
		Payload:  payload,
	})
}
