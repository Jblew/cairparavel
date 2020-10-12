package messengerapp

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/messengerapp/services"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// Bind to IoC container
func Bind(container *ioccontainer.Container) {
	container.Singleton(func(firestore *firestore.Client) domain.MessengerRecipientRepository {
		return &services.MessengerRecipientRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})
}
