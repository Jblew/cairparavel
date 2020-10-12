package domain

import "github.com/Jblew/ioccontainer/pkg/ioccontainer"

// MessengerReferral is a referral sent using messenger link to our app
type MessengerReferral struct {
	Code      string             `json:"code"`
	Recipient MessengerRecipient `json:"recipient"`
}

// OnNew handles new referral
func (referral *MessengerReferral) OnNew(container *ioccontainer.Container) error {
	var messengerRecipientRepository MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)
	var usersRepository UsersRepository
	container.Make(&usersRepository)

	userID := referral.Code
	err := messengerRecipientRepository.StoreForUser(userID, referral.Recipient)
	if err != nil {
		return err
	}
	user, err := usersRepository.GetUser(userID)
	if err != nil {
		return err
	}

	payload := make(map[string]interface{})
	payload["user"] = user
	return referral.Recipient.Notify(Notification{
		Template: "messenger_welcome",
		Payload:  payload,
	}, container)
}
