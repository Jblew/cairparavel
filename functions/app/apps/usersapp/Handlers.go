package usersapp

import "github.com/Jblew/cairparavel/functions/app/domain"

// OnUserCreated Handles a case when user is created
func OnUserCreated(user domain.User, usersRepository *domain.UsersRepository) error {
	return usersRepository.StoreUser(user)
}
