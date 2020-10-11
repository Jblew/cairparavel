package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// EventRepositoryFirestore implements EventRepository
type EventRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// GetEventByID retrives event by id
func (repo *EventRepositoryFirestore) GetEventByID(ID string) (domain.Event, error) {
	docRef := repo.Firestore.Collection(config.FirestorePaths.EventsCol()).Doc(ID)
	snapshot, err := docRef.Get(repo.Context)
	if err != nil {
		return domain.Event{}, err
	}

	if !snapshot.Exists() {
		return domain.Event{}, fmt.Errorf("No Event stored with ID=%s", ID)
	}

	var result domain.Event
	err = snapshot.DataTo(&result)
	if err != nil {
		return domain.Event{}, err
	}
	return result, nil
}
