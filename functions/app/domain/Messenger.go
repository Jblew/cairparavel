package domain

type MessengerRecipient struct {
	ID string `json:"id"`
}

type Messenger struct {
	Send(recipient MessengerRecipient, text string) error
}

type MessengerNotifier interface {
	SendNotification(recipient MessengerRecipient, notification Notification) error
}

type MessengerRecipientRepository interface {
	StoreMessengerRecipient(uid string, recipient MessengerRecipient) (User, error)
	GetMessengerRecipient(uid string) (MessengerRecipient, error)
}
