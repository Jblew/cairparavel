package notificationsapp

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/notificationsdomain"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/services"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	container.Singleton(func(firestore *firestore.Client) notificationsdomain.NotificationsRepository {
		return &services.NotificationsRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})

	container.Singleton(func() notificationsdomain.TemplatingService {
		return &services.TemplatingServiceGolang{}
	})

	container.Singleton(func() notificationsdomain.NotificationTemplateRepository {
		return &services.NotificationTemplateRepositoryStatic{}
	})
}
