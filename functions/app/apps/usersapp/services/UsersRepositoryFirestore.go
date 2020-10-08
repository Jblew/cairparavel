package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// UsersRepositoryFirestore implementation of domain.UsersService
type UsersRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   *context.Context
}

var usersColProd = "envs/prod/users"
var usersCols = []string{
	usersColProd,
	"envs/test/users",
}

// StoreUser saves user data
func (repo *UsersRepositoryFirestore) StoreUser(user domain.User) error {
	for _, usersCol := range usersCols {
		docRef := repo.Firestore.Collection(usersCol).Doc(user.UID)
		_, err := docRef.Create(*repo.Context, user)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUser returns user by id
func (repo *UsersRepositoryFirestore) GetUser(userID string) (domain.User, error) {
	docRef := repo.Firestore.Collection(usersColProd).Doc(userID)
	snapshot, err := docRef.Get(*repo.Context)
	if err != nil {
		return domain.User{}, err
	}

	if !snapshot.Exists() {
		return domain.User{}, fmt.Errorf("No such user with ID=%s", userID)
	}

	var result domain.User
	err = snapshot.DataTo(&result)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}
