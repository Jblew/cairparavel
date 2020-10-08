package services

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/notificationsdomain"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// NotificationsRepositoryFirestore implements NotificationsRepository
type NotificationsRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// AddNotificationToQueue adds notification to queue
func (repo *NotificationsRepositoryFirestore) AddNotificationToQueue(notification domain.Notification) error {
	_, err := repo.getNotificationsQueueColRef(notification.UID).NewDoc().Create(repo.Context, notification)
	if err != nil {
		return err
	}
	return nil
}

// AddNotificationToHistory adds notification to history
func (repo *NotificationsRepositoryFirestore) AddNotificationToHistory(notification notificationsdomain.PlainNotification) error {
	_, err := repo.getNotificationsHistoryColRef(notification.UID).NewDoc().Create(repo.Context, notification)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNotificationFromQueue deletes notification from queue
func (repo *NotificationsRepositoryFirestore) DeleteNotificationFromQueue(userID string, notificationID string) error {
	_, err := repo.getNotificationsQueueColRef(userID).Doc(notificationID).Delete(repo.Context)
	if err != nil {
		return err
	}
	return nil
}

func (repo *NotificationsRepositoryFirestore) getNotificationsQueueColRef(userID string) *firestore.CollectionRef {
	path := "/notifications/" + userID + "/queue"
	return repo.Firestore.Collection(path)
}

func (repo *NotificationsRepositoryFirestore) getNotificationsHistoryColRef(userID string) *firestore.CollectionRef {
	path := "/notifications/" + userID + "/history"
	return repo.Firestore.Collection(path)
}
