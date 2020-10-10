package functions

import (
	"context"
	"log"

	"github.com/Jblew/cairparavel/functions/app/domain"
)

// FnOnUserCreated handles user created event
func FnOnUserCreated(ctx context.Context, e AuthEvent) error {
	log.Printf("FnOnUserCreated: %v", e)

	user := domain.User{
		Email:    e.Email,
		UID:      e.UID,
		JoinedAt: e.Metadata.CreatedAt,
	}

	return user.OnAccountCreated(container)
}
