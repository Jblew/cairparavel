package messengerapp

import (
	"fmt"
	"log"

	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// OnMessengerInputMessage handles all input messages from messenger webhook
func OnMessengerInputMessage(message messenger.InputMessage, container *ioccontainer.Container) error {
	var lastErr error = nil
	for _, entry := range message.Entry {
		err := onInputEntry(entry, container)
		if err != nil {
			lastErr = err
			log.Fatal(err)
		}
	}
	return lastErr
}

func onInputEntry(entry messenger.InputMessageEntry, container *ioccontainer.Container) error {
	var lastErr error = nil
	for _, messagingEntry := range entry.Messaging {
		err := onMessagingEntry(messagingEntry, container)
		if err != nil {
			lastErr = err
			log.Fatal(err)
		}
	}
	return lastErr
}

func onMessagingEntry(entry messenger.MessagingEntry, container *ioccontainer.Container) error {
	var lastErr error = nil
	if entry.Referral != nil {
		err := onReferral(entry, entry.Referral, container)
		if err != nil {
			lastErr = err
			log.Fatal(err)
		}
	}
	if entry.Message != nil {
		err := onMessage(entry, entry.Message, container)
		if err != nil {
			lastErr = err
			log.Fatal(err)
		}
	}
	return lastErr
}

func onReferral(entry messenger.MessagingEntry, messagingReferral *messenger.MessagingEntryReferral, container *ioccontainer.Container) error {
	if len(messagingReferral.Ref) == 0 {
		return fmt.Errorf("Empty referral code")
	}
	referral := domain.MessengerReferral{
		Code: messagingReferral.Ref,
		Recipient: domain.MessengerRecipient{
			ID: entry.Sender.ID,
		},
	}

	return referral.OnNew(container)
}

func onMessage(entry messenger.MessagingEntry, messengerMessage *messenger.MessagingEntryMessage, container *ioccontainer.Container) error {
	message := domain.MessengerMessage{
		Text: messengerMessage.Text,
		Recipient: domain.MessengerRecipient{
			ID: entry.Sender.ID,
		},
	}
	return message.OnNew(container)
}
