package domain

import (
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// MessengerReferral is a referral sent using messenger link to our app
type MessengerReferral struct {
	Code      string             `json:"code" validate:"nonzero"`
	Recipient MessengerRecipient `json:"recipient" validate:"nonzero"`
}

// Validate validates
func (referral MessengerReferral) Validate() error {
	return validator.Validate(referral)
}

// OnNew handles new referral
func (referral *MessengerReferral) OnNew(container *ioccontainer.Container) error {
	err := referral.Validate()
	if err != nil {
		return err
	}

	var messengerRecipientRepository MessengerRecipientRepository
	container.Make(&messengerRecipientRepository)
	var usersRepository UsersRepository
	container.Make(&usersRepository)

	userID := referral.Code
	err = messengerRecipientRepository.StoreForUser(userID, referral.Recipient)
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
