package notificationsapp

import "github.com/Jblew/cairparavel/functions/app/domain"

// PlainNotification is notification resolved with template — used for history
type PlainNotification struct {
	Contents string
	UID      string
}

// NotificationTemplate is a template to parse notification into string message
type NotificationTemplate struct {
	Name     string
	Template string
}

// NotificationsRepository is a repository of notifications
type NotificationsRepository interface {
	AddNotificationToQueue(notification domain.Notification) error
	AddNotificationToHistory(notification PlainNotification) error
	DeleteNotificationFromQueue(notificationID string) error
}

// NotificationTemplateRepository is a repository of notification templates
type NotificationTemplateRepository interface {
	GetTemplate(Name string) (NotificationTemplate, error)
}

// TemplatingService is a service that parses the templates
type TemplatingService interface {
	ResolveTemplate(template string, payload map[string]interface{}) (string, error)
}
