package domain

import (
	"github.com/golobby/container/pkg/container"
)

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
