package messengerapp

import (
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	/*container.Singleton(func(firestore *firestore.Client) domain.MessengerRecipientRepository {
		return &services.MessengerRecipientRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})*/
}
