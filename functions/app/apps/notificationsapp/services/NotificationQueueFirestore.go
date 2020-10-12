package services

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// NotificationQueueFirestore implements NotificationQueue
type NotificationQueueFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// Add(uid string, notification Notification) error
// Delete(uid string, id string) error

// Add adds notification to queue
func (repo *NotificationQueueFirestore) Add(userID string, notification domain.Notification) error {
	notification.UID = userID

	path := config.FirestorePaths.MessengerNotificationsForUserCol(userID)
	docRef := repo.Firestore.Collection(path).NewDoc()

	notification.ID = docRef.ID
	_, err := docRef.Create(repo.Context, notification)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes notification from queue
func (repo *NotificationQueueFirestore) Delete(userID string, id string) error {
	collPath := config.FirestorePaths.MessengerNotificationsForUserCol(userID)
	_, err := repo.Firestore.Collection(collPath).Doc(id).Delete(repo.Context)
	if err != nil {
		return err
	}
	return nil
}
