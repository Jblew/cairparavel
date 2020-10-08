package domain

import "github.com/golobby/container/pkg/container"

func OnCommentAdded(event Event, comment EventComment, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["comment"] = comment

	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "comment_added",
		Payload:  payload,
	})
}

func OnEventStateChanged(event Event, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event

	if event.State == EventState.TIME_VOTING {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_voting_started",
			Payload:  payload,
		})
	} else if event.State == EventState.MEMBERS_SIGNUP {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_members_signup_started",
			Payload:  payload,
		})
	} else if event.State == EventState.SIGNUP_CLOSED {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_members_signup_closed",
			Payload:  payload,
		})
	} else if event.State == EventState.IN_PROGGRESS {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_started",
			Payload:  payload,
		})
	} else if event.State == EventState.CANCELLED {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_cancelled",
			Payload:  payload,
		})
	}
	return nil
}

func OnMessengerMessage(MessageText string, Recipient MessengerRecipient, container *container.Container) error {
	var messengerNotifier *MessengerNotifier
	container.Make(&MessengerNotifier)

	payload := make(map[string]interface{})
	payload["messageText"] = props.MessageText

	return messengerNotifier.SendNotification(props.Recipient, &Notification{
		Template: "messenger_respond",
		Payload:  payload,
	})
}

func OnMessengerReferral(ReferralCode string, MessengerRecipient MessengerRecipient, container *container.Container) error {
	var messengerRecipientRepository *MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)
	var messengerNotifier *MessengerNotifier
	container.Make(&messengerNotifier)
	var usersRepository *UsersRepository
	container.Make(&usersRepository)

	userID := props.ReferralCode
	err := messengerRecipientRepository.StoreMessengerUser(userID, props.MessengerRecipient)
	if err != nil {
		return err
	}
	user, err := usersRepository.GetUser(userID)
	if err != nil {
		return err
	}

	payload := make(map[string]interface{})
	payload["user"] = user
	return messengerNotifier.SendNotification(&Notification{
		Template: "messenger_welcome",
		Payload:  payload,
	})
}

func OnEventMemberSignup(event Event, signup EventSignup, econtainer *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["signup"] = signup
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "member_signed_in",
		Payload:  payload,
	})
}

func OnEventMemberSignout(event Event, signup EventSignup, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["signup"] = signup
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "member_signed_out",
		Payload:  payload,
	})
}

func OnEventVote(event Event, votes EventTimeVotes, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["votes"] = votes
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "event_voted",
		Payload:  payload,
	})
}

func OnEventVoteDeleted(event Event, votes EventTimeVotes, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	payload := make(map[string]interface{})
	payload["event"] = event
	payload["votes"] = votes
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "event_vote_deleted",
		Payload:  payload,
	})
}
