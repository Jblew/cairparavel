package usersapp

import (
	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/golobby/container/pkg/container"
)

// OnUserCreated Handles a case when user is created
func OnUserCreated(user domain.User, container container.Container) error {
	var usersRepository domain.UsersRepository
	container.Make(&usersRepository)

	return usersRepository.StoreUser(user)
}
