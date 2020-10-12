package domain

import (
	"fmt"
	"log"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// Notification is a notification to be sent to an user
type Notification struct {
	ID       string                 `json:"id"`
	UID      string                 `json:"uid" validate:"nonzero"`
	Template string                 `json:"template" validate:"nonzero"`
	Payload  map[string]interface{} `json:"payload" validate:"nonzero"`
}

// Validate validates
func (notification Notification) Validate(requireID bool) error {
	if requireID && len(notification.ID) == 0 {
		return fmt.Errorf("ID is required on Notification")
	}
	return validator.Validate(notification)
}

// OnQueued handles new notification in queue
func (notification *Notification) OnQueued(container *ioccontainer.Container) error {
	err := notification.Validate(true)
	if err != nil {
		return err
	}

	var queue NotificationQueue
	container.Make(&queue)

	var messengerRecipientRepository MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)

	log.Print("Loaded dependencies. Loading recipient...")
	recipient, err := messengerRecipientRepository.GetForUser(notification.UID)
	if err != nil {
		return err
	}
	log.Printf("Loaded recipient: %+v. Notifying...", recipient)

	err = recipient.Notify(*notification, container)
	if err != nil {
		return err
	}
	log.Printf("Notified recipient. Deleting...")

	return queue.Delete(notification.UID, notification.ID)
}

// NotificationQueue queues notifications
type NotificationQueue interface {
	Add(uid string, notification Notification) error
	Delete(uid string, id string) error
}
