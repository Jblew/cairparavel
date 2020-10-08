package messengerapp

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/messengerapp/services"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	container.Singleton(func(firestore *firestore.Client) domain.MessengerRecipientRepository {
		return &services.MessengerRecipientRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})
}
