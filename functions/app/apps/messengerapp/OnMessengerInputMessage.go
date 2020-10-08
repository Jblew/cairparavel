package messengerapp

import (
	"fmt"
	"log"

	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
	"github.com/golobby/container/pkg/container"
)

// OnMessengerInputMessage handles all input messages from messenger webhook
func OnMessengerInputMessage(message messenger.InputMessage, container container.Container) error {
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

func onInputEntry(entry messenger.InputMessageEntry, container container.Container) error {
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

func onMessagingEntry(entry messenger.MessagingEntry, container container.Container) error {
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

func onReferral(entry messenger.MessagingEntry, referral *messenger.MessagingEntryReferral, container container.Container) error {
	if len(referral.Ref) == 0 {
		return fmt.Errorf("Empty referral code")
	}
	return domain.OnMessengerReferral(referral.Ref, domain.MessengerRecipient{
		ID: entry.Sender.ID,
	}, container)
}

func onMessage(entry messenger.MessagingEntry, message *messenger.MessagingEntryMessage, container container.Container) error {
	return domain.OnMessengerMessage(message.Text, domain.MessengerRecipient{
		ID: entry.Sender.ID,
	}, container)
}
