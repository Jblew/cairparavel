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

type MessengerIDsRepository interface {
	StoreMessengerUser(uid string, recipient MessengerRecipient) (User, error)
}
