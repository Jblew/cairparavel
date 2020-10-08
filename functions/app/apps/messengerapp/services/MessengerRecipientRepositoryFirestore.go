package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// MessengerRecipientRepositoryFirestore implements MessengerRecipientRepository
type MessengerRecipientRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

var colPath string = "notification_messengerid"

// StoreMessengerRecipient saves messenger recipient id
func (repo *MessengerRecipientRepositoryFirestore) StoreMessengerRecipient(uid string, recipient domain.MessengerRecipient) error {
	docRef := repo.Firestore.Collection(colPath).Doc(uid)
	_, err := docRef.Set(repo.Context, recipient)
	return err
}

// GetMessengerRecipient retrives messenger recipient id
func (repo *MessengerRecipientRepositoryFirestore) GetMessengerRecipient(uid string) (domain.MessengerRecipient, error) {
	docRef := repo.Firestore.Collection(colPath).Doc(uid)
	snapshot, err := docRef.Get(repo.Context)
	if err != nil {
		return domain.MessengerRecipient{}, err
	}

	if !snapshot.Exists() {
		return domain.MessengerRecipient{}, fmt.Errorf("No MessengerRecipient stored for user with UID=%s", uid)
	}

	var result domain.MessengerRecipient
	err = snapshot.DataTo(&result)
	if err != nil {
		return domain.MessengerRecipient{}, err
	}
	return result, nil
}
