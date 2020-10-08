package domain

import "time"

type User struct {
	Email    string    `json:"Email"`
	UID      string    `json:"UID"`
	JoinedAt time.Time `json:"JoinedAt"`
}

type UsersRepository interface {
	StoreUser(user User) error
}
