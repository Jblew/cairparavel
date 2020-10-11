package domain

import "github.com/golobby/container/pkg/container"

// MessengerRecipient message recipient or sender in FB messenger
type MessengerRecipient struct {
	ID string `json:"id"`
}

// Notify sends notification to messenger user
func (recipient *MessengerRecipient) Notify(notification Notification, container container.Container) error {
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
