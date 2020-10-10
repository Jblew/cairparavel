package domain

// Notification is a notification to be sent to an user
type Notification struct {
	ID       string                 `json:"id"`
	UID      string                 `json:"uid"`
	Template string                 `json:"template"`
	Payload  map[string]interface{} `json:"payload"`
}

// NotificationQueue queues notifications
type NotificationQueue interface {
	ScheduleToSend(uid string, notification Notification) error
}

type EventObserversNotifier interface {
	NotifyEventObservers(event Event, notification Notification) error
}
