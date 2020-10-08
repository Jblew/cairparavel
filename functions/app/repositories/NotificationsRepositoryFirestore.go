import 	(
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

type NotificationsRepositoryFirestore struct {
	Firestore *firestore.Client
	Context *context.Context
}

func (repo *NotificationsRepositoryFirestore) AddNotificationToQueue(notification domain.Notification) error {
	_, err := repo.getNotificationsQueueColRef(notification.UID).Create(notification)
	if err != nil {
		return err
	}
	return nil
}

func (repo *NotificationsRepositoryFirestore) AddNotificationToHistory(notification PlainNotification) error {
	_, err := repo.getNotificationsHistoryColRef(notification.UID).Create(notification)
	if err != nil {
		return err
	}
	return nil
}

func (repo *NotificationsRepositoryFirestore) DeleteNotificationFromQueue(notificationID string) error {
	_, err := repo.getNotificationsQueueColRef(notification.UID).Doc(notification).Delete()
	if err != nil {
		return err
	}
	return nil
}


func (UsersRepositoryFirestore *repo) getNotificationsQueueColRef(userId string) *firestore.CollectionRef {
	path := fmt.Printf("/notifications/%s/queue", userId)
	return repo.Firestore.Collection(path)
}

func (UsersRepositoryFirestore *repo) getNotificationsHistoryColRef(userId string) *firestore.CollectionRef {
	path := fmt.Printf("/notifications/%s/history", userId)
	return repo.Firestore.Collection(path)
}
