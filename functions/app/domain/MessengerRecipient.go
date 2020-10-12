package domain

import (
	"log"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// MessengerRecipient message recipient or sender in FB messenger
type MessengerRecipient struct {
	ID string `json:"id" validate:"nonzero"`
}

// Validate validates
func (recipient MessengerRecipient) Validate() error {
	return validator.Validate(recipient)
}

// Notify sends notification to messenger user
func (recipient *MessengerRecipient) Notify(notification Notification, container *ioccontainer.Container) error {
	err := recipient.Validate()
	if err != nil {
		return err
	}

	var service MessengerNotificationService
	container.Make(&service)

	log.Printf("Sending notification to messenger recipient %s: %+v", recipient.ID, notification)
	return service.SendNotification(recipient.ID, notification.Template, notification.Payload)
}

// MessengerRecipientRepository stores or retrives FB messenger recipient ID based on our UID
type MessengerRecipientRepository interface {
	StoreForUser(uid string, recipient MessengerRecipient) error
	GetForUser(uid string) (MessengerRecipient, error)
}

// MessengerNotificationService sends notification to messenger user
type MessengerNotificationService interface {
	SendNotification(recipientID string, templateName string, payload map[string]interface{}) error
}
