package notificationsapp

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/notificationsdomain"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/services"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	container.Singleton(func(firestore *firestore.Client) domain.NotificationQueue {
		return &services.NotificationQueueFirestore{
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

	container.Singleton(func(
		messenger messenger.Messenger,
		templatingRepo notificationsdomain.NotificationTemplateRepository,
		templatingService notificationsdomain.TemplatingService,
	) domain.MessengerNotificationService {
		return &services.MessengerNotificationService{
			Messenger:          messenger,
			TemplateRepository: templatingRepo,
			TemplatingService:  templatingService,
		}
	})
}
