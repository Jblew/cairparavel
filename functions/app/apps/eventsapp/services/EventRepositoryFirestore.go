package services

import (
	"context"
	"fmt"
	"log"
	"time"

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

	result, err := eventFromSnapshot(snapshot, false)
	if err != nil {
		return domain.Event{}, err
	}
	result.ID = snapshot.Ref.ID
	return result, nil
}

// GetAllNonFinishedAt retrives events that are not yet finished
func (repo *EventRepositoryFirestore) GetAllNonFinishedAt(atTime time.Time) ([]domain.Event, error) {
	colRef := repo.Firestore.Collection(config.FirestorePaths.EventsCol())
	nowMillis := atTime.UnixNano() / int64(time.Millisecond)
	snapshots, err := colRef.Where("endTime", ">", nowMillis).Documents(repo.Context).GetAll()
	if err != nil {
		return []domain.Event{}, err
	}
	results := make([]domain.Event, 0, len(snapshots))

	for _, snapshot := range snapshots {
		event, err := eventFromSnapshot(snapshot, false)
		if err != nil {
			log.Printf("Invalid event fetched: %v", err)
		} else {
			event.ID = snapshot.Ref.ID
			results = append(results, event)
		}
	}
	return results, nil
}

func eventFromSnapshot(snapshot *firestore.DocumentSnapshot, requireID bool) (domain.Event, error) {
	var event domain.Event
	err := snapshot.DataTo(&event)
	if err != nil {
		return domain.Event{}, err
	}
	if err = event.Validate(requireID); err != nil {
		return domain.Event{}, err
	}
	return event, nil
}
