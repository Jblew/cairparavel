package domain

import "github.com/golobby/container/pkg/container"

// MessengerMessage is message send from messenger to our bot
type MessengerMessage struct {
	Text      string `"json":"text"`
	Recipient MessengerRecipient
}

// OnNew handles new message
func (message *MessengerMessage) OnNew(container container.Container) error {
	notification := Notification{
		Template: "messenger_respond",
		Payload:  make(map[string]interface{}),
	}
	notification.Payload["messageText"] = message.Text

	return message.Recipient.Notify(notification, container)
}
