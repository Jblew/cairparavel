package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/google/uuid"
)

// NotificationQueueFirestore implements NotificationQueue
type NotificationQueueFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// Add adds notification to queue
func (repo *NotificationQueueFirestore) Add(userID string, notification domain.Notification) error {
	if len(userID) == 0 {
		return fmt.Errorf("Empty userID")
	}

	notification.UID = userID
	if err := notification.Validate(false); err != nil {
		return err
	}

	path := config.FirestorePaths.MessengerNotificationsForUserCol(userID)
	log.Printf("NotificationQueueFirestore->notification.ID = docRef.ID")
	notification.ID = uuid.New().String()

	log.Printf("NotificationQueueFirestore->Collection(" + path + ")")
	colRef := repo.Firestore.Collection(path)
	log.Printf("NotificationQueueFirestore->colRef.Doc(notification.ID)")
	docRef := colRef.Doc(notification.ID)

	log.Printf("NotificationQueueFirestore->docRef.Create")
	row, err := notificationToRow(notification)
	if err != nil {
		return err
	}
	_, err = docRef.Create(repo.Context, row)
	if err != nil {
		log.Printf("NotificationQueueFirestore->docRef.Create done with error: %v", err)
		return err
	}
	log.Printf("NotificationQueueFirestore->docRef.Create done without error")
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

type notificationRow struct {
	ID       string
	UID      string
	Template string
	Payload  string
}

func (row *notificationRow) toNotification() (domain.Notification, error) {
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(row.Payload), &payload)
	if err != nil {
		return domain.Notification{}, err
	}

	return domain.Notification{
		ID:       row.ID,
		UID:      row.UID,
		Template: row.Template,
		Payload:  payload,
	}, nil
}

func notificationToRow(notification domain.Notification) (notificationRow, error) {
	payloadOut, err := json.Marshal(notification.Payload)
	if err != nil {
		return notificationRow{}, err
	}

	return notificationRow{
		ID:       notification.ID,
		UID:      notification.UID,
		Template: notification.Template,
		Payload:  string(payloadOut),
	}, nil
}
