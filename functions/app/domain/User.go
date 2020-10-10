package domain

import (
	"time"

	"github.com/golobby/container/pkg/container"
)

// User stored in db. Do not confuse with firestore user
type User struct {
	Email    string    `json:"Email"`
	UID      string    `json:"UID"`
	JoinedAt time.Time `json:"JoinedAt"`
}

// OnAccountCreated handles situation when user logs in and doesnt have an account stored
func (user *User) OnAccountCreated(temporaryUser User, container container.Container) error {
	var usersRepository UsersRepository
	container.Make(&usersRepository)

	return usersRepository.StoreUser(*user)
}

// UsersRepository repository for users
type UsersRepository interface {
	StoreUser(user User) error
	GetUser(userID string) (User, error)
}
