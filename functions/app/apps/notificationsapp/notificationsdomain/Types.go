package notificationsdomain

// NotificationTemplate is a template to parse notification into string message
type NotificationTemplate struct {
	Name     string
	Template string
}

// NotificationTemplateRepository is a repository of notification templates
type NotificationTemplateRepository interface {
	GetTemplate(Name string) (NotificationTemplate, error)
}

// TemplatingService is a service that parses the templates
type TemplatingService interface {
	ResolveTemplate(template string, payload map[string]interface{}) (string, error)
}
