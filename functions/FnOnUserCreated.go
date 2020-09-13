package functions

import (
	"context"
	"log"
	"time"
)

// AuthEvent is the payload of a Firestore Auth event.
type AuthEvent struct {
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	UID string `json:"uid"`
}

// UserDoc is type for /users/{uid} document
type UserDoc struct {
	Email    string    `json:"email"`
	UID      string    `json:"uid"`
	JoinedAt time.Time `json:"joinedAt"`
}

const usersCol = "users"

// FnOnUserCreated handles user created event
func FnOnUserCreated(ctx context.Context, e AuthEvent) error {
	printDebug(e)

	userDoc := buildUserDoc(e)
	return publishUserDoc(userDoc)
}

func buildUserDoc(e AuthEvent) UserDoc {
	return UserDoc{
		Email:    e.Email,
		UID:      e.UID,
		JoinedAt: e.Metadata.CreatedAt,
	}
}

func publishUserDoc(userDoc UserDoc) error {
	docRef := application.Firestore.Collection(usersCol).Doc(userDoc.UID)
	_, err := docRef.Create(application.Context, userDoc)
	return err
}

func printDebug(e AuthEvent) {
	log.Printf("FnOnUserCreated: %v", e)
}
