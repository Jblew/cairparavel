package domain

import (
	"log"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// MessengerMessage is message send from messenger to our bot
type MessengerMessage struct {
	Text      string `json:"text" validate:"nonzero"`
	Recipient MessengerRecipient
}

// Validate validates
func (message MessengerMessage) Validate() error {
	return validator.Validate(message)
}

// OnNew handles new message
func (message *MessengerMessage) OnNew(container *ioccontainer.Container) error {
	log.Printf("MessengerMessage.OnNew: %+v", message)
	/*err := message.Validate()
	if err != nil {
		return err
	}

	notification := Notification{
		Template: "messenger_respond",
		Payload:  make(map[string]interface{}),
	}
	notification.Payload["messageText"] = message.Text

	return message.Recipient.Notify(notification, container)*/
	// Remember that we get notified about our messages too. There is a possibility to generate a message loop
	return nil
}
