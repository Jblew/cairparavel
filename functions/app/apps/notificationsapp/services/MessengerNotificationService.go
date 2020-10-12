package services

import (
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/notificationsdomain"
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
)

// MessengerNotificationService implements domain.MessengerNotificationService
type MessengerNotificationService struct {
	TemplateRepository notificationsdomain.NotificationTemplateRepository
	TemplatingService  notificationsdomain.TemplatingService
	Messenger          messenger.Messenger
}

// SendNotification sends notification to messenger user
func (service *MessengerNotificationService) SendNotification(recipientID string, templateName string, payload map[string]interface{}) error {
	template, err := service.TemplateRepository.GetTemplate(templateName)
	if err != nil {
		return err
	}

	messageStr, err := service.TemplatingService.ResolveTemplate(template.Template, payload)
	if err != nil {
		return err
	}

	return service.Messenger.SendMessage(messenger.Recipient{
		ID: recipientID,
	}, messageStr)
}
