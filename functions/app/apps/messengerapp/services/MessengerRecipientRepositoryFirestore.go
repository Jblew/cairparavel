package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// MessengerRecipientRepositoryFirestore implements MessengerRecipientRepository
type MessengerRecipientRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// StoreForUser saves messenger recipient id
func (repo *MessengerRecipientRepositoryFirestore) StoreForUser(uid string, recipient domain.MessengerRecipient) error {
	if err := recipient.Validate(); err != nil {
		return err
	}
	docRef := repo.Firestore.Doc(config.FirestorePaths.MessengerRecipientForUserDoc(uid))
	_, err := docRef.Set(repo.Context, recipient)
	return err
}

// GetForUser retrives messenger recipient id
func (repo *MessengerRecipientRepositoryFirestore) GetForUser(uid string) (domain.MessengerRecipient, error) {
	docRef := repo.Firestore.Doc(config.FirestorePaths.MessengerRecipientForUserDoc(uid))
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
	if err := result.Validate(); err != nil {
		return domain.MessengerRecipient{}, err
	}
	return result, nil
}
