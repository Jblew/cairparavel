package domain

import (
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

	return service.SendNotification(recipient.ID, notification)
}

// MessengerRecipientRepository stores or retrives FB messenger recipient ID based on our UID
type MessengerRecipientRepository interface {
	StoreForUser(uid string, recipient MessengerRecipient) error
	GetForUser(uid string) (MessengerRecipient, error)
}

// MessengerNotificationService sends notification to messenger user
type MessengerNotificationService interface {
	SendNotification(recipientID string, notification Notification) error
}
