package functions

import (
	"context"
	"log"
	"time"

	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/cairparavel/functions/util"
)

// FnOnFirebaseAccountCreated handles user created event
func FnOnFirebaseAccountCreated(ctx context.Context, e AuthEvent) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnFirebaseAccountCreated",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		log.Printf("FnOnFirebaseAccountCreated: %v", e)

		user := domain.User{
			Email:    e.Email,
			UID:      e.UID,
			JoinedAt: e.Metadata.CreatedAt,
		}

		return user.OnAccountCreated(container)
	})
}

// AuthEvent is the payload of a Firestore Auth event.
type AuthEvent struct {
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	UID string `json:"uid"`
}
