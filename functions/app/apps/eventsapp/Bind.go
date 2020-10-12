package eventsapp

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/eventsapp/services"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	container.Singleton(func(firestore *firestore.Client) domain.EventRepository {
		return &services.EventRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})

	container.Singleton(func(firestore *firestore.Client) domain.EventObserverRepository {
		return &services.EventObserverRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})

	container.Singleton(func(firestore *firestore.Client) domain.EventSignupRepository {
		return &services.EventSignupRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})

	container.Singleton(func(firestore *firestore.Client) EventStateChangedRepository {
		return &services.EventStateChangedRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})
}
