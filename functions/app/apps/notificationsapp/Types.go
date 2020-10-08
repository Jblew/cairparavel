package notificationsapp

import "github.com/Jblew/cairparavel/functions/app/domain"

type PlainNotification struct {
	Contents string
	UID      sring
}

type NotificationTemplate {
	Name string
	Template string
}

type NotificationsRepository interface {
	AddNotificationToQueue(notification domain.Notification) error
	AddNotificationToHistory(notification PlainNotification) error
	DeleteNotificationFromQueue(notificationID string) error
}

type NotificationTemplateRepository interface {
	GetTemplate(Name string) (NotificationTemplate, error)
}

type TemplatingService interface {
	ResolveTemplate(template string, payload map[string]interface{}) (string, error)
}
