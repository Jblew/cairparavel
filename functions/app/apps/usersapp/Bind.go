package usersapp

import (
	firestore "cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/apps/usersapp/services"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC container
func Bind(container container.Container) {
	container.Singleton(func(firestore *firestore.Client) domain.UsersRepository {
		return &services.UsersRepositoryFirestore{
			Firestore: firestore,
		}
	})
}
