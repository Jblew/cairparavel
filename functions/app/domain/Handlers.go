package domain


func OnCommentAdded(event Event, comment EventComment, eventObserversNotifier *EventObserversNotifier) error {
	notification := {
		Template: "comment_added",
		Payload: {
			Event: event,
			Comment: comment,
		}
	}
	return eventObserversNotifier.NotifyEventObservers(event, notification)
}

func OnEventStateChanged(event Event, eventObserversNotifier *EventObserversNotifier) error {
	if event.State == EventState.TIME_VOTING {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_voting_started",
			Payload: {
				Event: event,
			}
		})
	}
	else if event.State == EventState.MEMBERS_SIGNUP {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_members_signup_started",
			Payload: {
				Event: event,
			}
		})
	}
	else if event.State == EventState.SIGNUP_CLOSED {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_members_signup_closed",
			Payload: {
				Event: event,
			}
		})
	}
	else if event.State == EventState.IN_PROGGRESS {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_started",
			Payload: {
				Event: event,
			}
		})
	}
	else if event.State == EventState.CANCELLED {
		return eventObserversNotifier.NotifyEventObservers(event, &Notification{
			Template: "event_cancelled",
			Payload: {
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
func OnMessengerMessage(props OnMessengerMessageProps) error {
	return props.MessengerNotifier.SendNotification(props.Recipient, &Notification{
		Template: "messenger_respond",
		Payload: {
			MessageText props.MessageText,
		}
	})
}

type OnMessengerReferralProps struct {
	ReferralCode string,
	MessengerRecipient: MessengerRecipient,
	MessengerNotifier *MessengerNotifier
	MessengerIDsRepository *MessengerIDsRepository
}
func OnMessengerReferral(props OnMessengerReferralProps) error {
	user, err := MessengerIDsRepository.StoreMessengerUser(props.ReferralCode, props.MessengerRecipient)
	if err != nil {
		return err
	}
	returnprops.MessengerNotifier.SendNotification(&Notification{
		Template: "messenger_welcome",
		Payload: {
			User: user,
		}
	})
}

func OnEventMemberSignup(event Event, signup EventSignup, eventObserversNotifier *EventObserversNotifier) error {
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "member_signed_in",
		Payload: {
			Event: event,
			Signup: signup,
		}
	})
}

func OnEventMemberSignout(event Event, signup EventSignup, eventObserversNotifier *EventObserversNotifier) error {
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "member_signed_out",
		Payload: {
			Event: event,
			Signup: signup,
		}
	})
}

func OnEventVote(event Event, votes EventTimeVotes, eventObserversNotifier *EventObserversNotifier) error {
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "event_voted",
		Payload: {
			Event: event,
			Votes : votes,
		}
	})
}

func OnEventVoteDeleted(event Event, votes EventTimeVotes, eventObserversNotifier *EventObserversNotifier) error {
	return eventObserversNotifier.NotifyEventObservers(event, &Notification{
		Template: "event_vote_deleted",
		Payload: {
			Event: event,
			Votes: votes,
		}
	})
}
