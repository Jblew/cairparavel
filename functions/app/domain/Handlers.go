package domain

import "github.com/golobby/container/pkg/container"

func OnCommentAdded(event Event, comment EventComment, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "comment_added",
			Payload: struct{
				Event: event,
				Comment: comment,
			}
	})
}

func OnEventStateChanged(event Event, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	if event.State == EventState.TIME_VOTING {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_voting_started",
			Payload: struct {
				Event: event,
			}
		})
	}
	else if event.State == EventState.MEMBERS_SIGNUP {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_members_signup_started",
			Payload: struct {
				Event: event,
			}
		})
	}
	else if event.State == EventState.SIGNUP_CLOSED {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_members_signup_closed",
			Payload: struct {
				Event: event,
			}
		})
	}
	else if event.State == EventState.IN_PROGGRESS {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_started",
			Payload: struct {
				Event: event,
			}
		})
	}
	else if event.State == EventState.CANCELLED {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_cancelled",
			Payload: struct {
				Event: event,
			}
		})
	}
	return nil
}

type OnMessengerMessageProps struct {
	MessageText string,
	Recipient MessengerRecipient,
	MessengerNotifier *MessengerNotifier,
}
func OnMessengerMessage(MessageText string, Recipient MessengerRecipient, container *container.Container) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	return props.MessengerNotifier.SendNotification(props.Recipient, &Notification{
		Template: "messenger_respond",
		Payload: struct {
			MessageText props.MessageText,
		}
	})
}

type OnMessengerReferralProps struct {
	ReferralCode string,
	MessengerRecipient: MessengerRecipient,
	MessengerNotifier *MessengerNotifier,
	MessengerRecipientRepository *MessengerRecipientRepository,
	UsersRepository *UsersRepository,
}
func OnMessengerReferral(props OnMessengerReferralProps) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	userID := props.ReferralCode
	err := props.MessengerRecipientRepository.StoreMessengerUser(userID, props.MessengerRecipient)
	if err != nil {
		return err
	}
	user, err := props.UsersRepository.GetUser(userID)
	if err != nil {
		return err
	}

	returnprops.MessengerNotifier.SendNotification(&Notification{
		Template: "messenger_welcome",
		Payload: struct {
			User: user,
		}
	})
}

func OnEventMemberSignup(event Event, signup EventSignup, eventObserversNotifier *EventObserversNotifier) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "member_signed_in",
		Payload: struct {
			Event: event,
			Signup: signup,
		}
	})
}

func OnEventMemberSignout(event Event, signup EventSignup, eventObserversNotifier *EventObserversNotifier) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "member_signed_out",
		Payload: struct {
			Event: event,
			Signup: signup,
		}
	})
}

func OnEventVote(event Event, votes EventTimeVotes, eventObserversNotifier *EventObserversNotifier) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "event_voted",
		Payload: struct {
			Event: event,
			Votes : votes,
		}
	})
}

func OnEventVoteDeleted(event Event, votes EventTimeVotes, eventObserversNotifier *EventObserversNotifier) error {
	var eventObserversNotifier *EventObserversNotifier
	container.Make(&eventObserversNotifier)

	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "event_vote_deleted",
		Payload: struct {
			Event: event,
			Votes: votes,
		}
	})
}
