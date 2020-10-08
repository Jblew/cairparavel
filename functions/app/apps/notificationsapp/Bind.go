package notificationsapp

import (
	firestore "cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/services"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	container.Singleton(func(firestore *firestore.Client) *NotificationsRepository {
		return &services.NotificationsRepositoryFirestore{
			Firestore: firestore,
		}
	})

	container.Singleton(func() TemplatingService {
		return &services.TemplatingServiceGolang{}
	})

	container.Singleton(func() NotificationTemplateRepository {
		return &services.NotificationTemplateRepositoryStatic{}
	})
}
