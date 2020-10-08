package notificationsapp

import (
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/golobby/container/pkg/container"
)

// OnNotificationCreated Handles a case when notification is awaiting sending
func OnNotificationCreated(notification domain.Notification, container *container.Container) error {
	var templateRepository NotificationTemplateRepository
	container.Make(&templateRepository)

	var templatingService TemplatingService
	container.Make(&templatingService)

	var notificationsRepository NotificationsRepository
	container.Make(&notificationsRepository)

	var messengerRecipientRepository domain.MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)

	var messengerNotifier domain.MessengerNotifier
	container.Make(&messengerNotifier)

	template, err := templateRepository.GetTemplate(notification.Template)
	if err != nil {
		return err
	}
	messageText, err := templatingService.ResolveTemplate(template.Template, notification.Payload)
	if err != nil {
		return err
	}

	err = notificationsRepository.AddNotificationToHistory(PlainNotification{
		Contents: messageText,
		UID:      notification.UID,
	})
	if err != nil {
		return err
	}

	recipient, err := messengerRecipientRepository.GetMessengerRecipient(notification.UID)
	if err != nil {
		return err
	}

	err = messengerNotifier.SendNotification(recipient, notification)
	if err != nil {
		return err
	}

	return notificationsRepository.DeleteNotificationFromQueue(notification.ID)
}
