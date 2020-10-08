package domain

type Notification struct {
	Template: string `json:"template"`,
	Payload: map[string]interface{}
}

type EventObserversNotifier interface {
	NotifyEventObservers(event Event, notification Notification) error
}


type NotificationTemplateEngine interface {
	ParseNotification(notification Notification) (string, error)
}
