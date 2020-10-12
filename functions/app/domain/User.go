package domain

import (
	"time"

	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
	"gopkg.in/validator.v2"
)

// User stored in db. Do not confuse with firestore user
type User struct {
	Email    string    `json:"Email" validate:"nonzero"`
	UID      string    `json:"UID" validate:"nonzero"`
	JoinedAt time.Time `json:"JoinedAt"`
}

// Validate validates
func (user User) Validate() error {
	return validator.Validate(user)
}

// OnAccountCreated handles situation when user logs in and doesnt have an account stored
func (user *User) OnAccountCreated(container *ioccontainer.Container) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	var usersRepository UsersRepository
	container.Make(&usersRepository)

	return usersRepository.StoreUser(*user)
}

// UsersRepository repository for users
type UsersRepository interface {
	StoreUser(user User) error
	GetUser(userID string) (User, error)
}
