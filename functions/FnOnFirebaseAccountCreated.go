package functions

import (
	"context"
	"log"
	"time"

	"github.com/Jblew/cairparavel/functions/app/domain"
)

// FnOnFirebaseAccountCreated handles user created event
func FnOnFirebaseAccountCreated(ctx context.Context, e AuthEvent) error {
	log.Printf("FnOnFirebaseAccountCreated: %v", e)

	user := domain.User{
		Email:    e.Email,
		UID:      e.UID,
		JoinedAt: e.Metadata.CreatedAt,
	}

	return user.OnAccountCreated(container)
}

// AuthEvent is the payload of a Firestore Auth event.
type AuthEvent struct {
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	UID string `json:"uid"`
}
