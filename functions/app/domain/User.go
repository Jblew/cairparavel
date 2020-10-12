package domain

import (
	"time"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// User stored in db. Do not confuse with firestore user
type User struct {
	Email    string    `json:"Email"`
	UID      string    `json:"UID"`
	JoinedAt time.Time `json:"JoinedAt"`
}

// OnAccountCreated handles situation when user logs in and doesnt have an account stored
func (temporaryUser *User) OnAccountCreated(container *ioccontainer.Container) error {
	var usersRepository UsersRepository
	container.Make(&usersRepository)

	return usersRepository.StoreUser(*temporaryUser)
}

// UsersRepository repository for users
type UsersRepository interface {
	StoreUser(user User) error
	GetUser(userID string) (User, error)
}
