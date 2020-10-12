package usersapp

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/usersapp/services"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// Bind to IoC container
func Bind(container *ioccontainer.Container) {
	container.Singleton(func(firestore *firestore.Client) domain.UsersRepository {
		return &services.UsersRepositoryFirestore{
			Firestore: firestore,
			Context:   context.Background(),
		}
	})
}
