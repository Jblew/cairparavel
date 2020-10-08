package domain

// MessengerRecipient message recipient or sender in FB messenger
type MessengerRecipient struct {
	ID string `json:"id"`
}

// MessengerNotifier sends notification to messenger user
type MessengerNotifier interface {
	SendNotification(recipient MessengerRecipient, notification Notification) error
}

// MessengerRecipientRepository stores or retrives FB messenger recipient ID based on our UID
type MessengerRecipientRepository interface {
	StoreMessengerRecipient(uid string, recipient MessengerRecipient) error
	GetMessengerRecipient(uid string) (MessengerRecipient, error)
}
