package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// UsersRepositoryFirestore implementation of domain.UsersService
type UsersRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

var usersCols = []string{
	config.FirestorePaths.ProdUsersCol(),
	config.FirestorePaths.TestUsersCol(),
}

// StoreUser saves user data
func (repo *UsersRepositoryFirestore) StoreUser(user domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	for _, usersCol := range usersCols {
		docRef := repo.Firestore.Collection(usersCol).Doc(user.UID)
		_, err := docRef.Create(repo.Context, user)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUser returns user by id
func (repo *UsersRepositoryFirestore) GetUser(userID string) (domain.User, error) {
	docRef := repo.Firestore.Doc(config.FirestorePaths.ProdUserDoc(userID))
	snapshot, err := docRef.Get(repo.Context)
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
	if err := result.Validate(); err != nil {
		return domain.User{}, err
	}
	return result, nil
}
