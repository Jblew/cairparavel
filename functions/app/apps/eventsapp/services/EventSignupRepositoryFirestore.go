package services

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
)

// EventSignupRepositoryFirestore implements EventSignupRepository
type EventSignupRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// GetCount retrives count of signups
func (repo *EventSignupRepositoryFirestore) GetCount(eventID string) (int, error) {
	collectionPath := config.FirestorePaths.SignupsForEventCol(eventID)
	collectionRef := repo.Firestore.Collection(collectionPath)
	docRef, err := collectionRef.DocumentRefs(repo.Context).GetAll()
	if err != nil {
		return 0, err
	}

	return len(docRef), nil
}
