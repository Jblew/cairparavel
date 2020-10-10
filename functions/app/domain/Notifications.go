package domain

import (
	"github.com/golobby/container/pkg/container"
)

// Notification is a notification to be sent to an user
type Notification struct {
	ID       string                 `json:"id"`
	UID      string                 `json:"uid"`
	Template string                 `json:"template"`
	Payload  map[string]interface{} `json:"payload"`
}

// PlainNotification is notification resolved with template — used for history
type PlainNotification struct {
	Contents string
	UID      string
}

// OnQueued handles new notification in queue
func (notification *Notification) OnQueued(container container.Container) error {
	var queue NotificationQueue
	container.Make(&queue)

	var messengerRecipientRepository MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)

	recipient, err := messengerRecipientRepository.GetForUser(notification.UID)
	if err != nil {
		return err
	}

	err = recipient.Notify(*notification, container)
	if err != nil {
		return err
	}

	return queue.Delete(notification.UID, notification.ID)
}

// NotificationQueue queues notifications
type NotificationQueue interface {
	Add(uid string, notification Notification) error
	Delete(uid string, id string) error
}
