package usersapp

import "github.com/Jblew/cairparavel/functions/app/domain"

type OnNotificationCreatedProps struct {
	Notification *domain.Notification,
	TemplateRepository *NotificationTemplateRepository,
	TemplatingService *TemplatingService,
	NotificationsRepository *NotificationsRepository,
	MessengerNotifier *domain.MessengerNotifier,
	MessengerRecipientRepository *domain.MessengerRecipientRepository,
}

// OnNotificationCreated Handles a case when notification is awaiting sending
func OnNotificationCreated(props OnNotificationCreatedProps) error {
	template, err := props.TemplateRepository.GetTemplate(props.Notification.Template)
	if err != nil {
		return err
	}
	messageText, err := props.TemplatingService.ResolveTemplate(template, props.Notification.Payload)
	if err != nil {
		return err
	}

	err = props.NotificationsRepository.AddNotificationToHistory(&PlainNotification{
			Contents: messageText,
			UID: props.Notification.UID,
	})
	if err != nil {
		return err
	}

	recipient, err := props.MessengerRecipientRepository.GetMessengerRecipient(props.Notification.UID)
	if err != nil {
		return err
	}

	err = props.MessengerNotifier.SendNotification(recipient, props.Notification)
	if err != nil {
		return err
	}

	return props.NotificationsRepository.DeleteNotificationFromQueue(props.Notification.ID)
}
