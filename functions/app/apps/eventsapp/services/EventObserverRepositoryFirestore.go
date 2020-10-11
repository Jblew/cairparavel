package services

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// EventObserverRepositoryFirestore implements EventObserverRepository
type EventObserverRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

type EventObserverRepository interface {
	GetAllForEvent(eventID string) ([]EventObserver, error)
	Add(observer EventObserver) error
}

// GetAllForEvent retrives observers of an event
func (repo *EventObserverRepositoryFirestore) GetAllForEvent(eventID string) ([]domain.EventObserver, error) {
	collectionRef := repo.Firestore.Collection(config.FirestorePaths.ObserversForEventCol(eventID))
	snapshots, err := collectionRef.GetAll(repo.Context)
	if err != nil {
		return []domain.EventObserver{}, err
	}

	results := make([]domain.EventObserver, len(snapshots))

	for _, snapshot := range snapshots {
		var observer domain.EventObserver
		err = snapshot.DataTo(&observer)
		if err != nil {
			return []domain.EventObserver, err
		}
		results = append(results, observer)
	}
	return results, nil
}

// Add saves event observer
func (repo *UsersRepositoryFirestore) Add(observer EventObserver) error {
	docRef := repo.Firestore.Doc(config.FirestorePaths.ObserversForEventForUserDoc(observer.EventID, observer.UID))
	_, err := docRef.Create(repo.Context, observer)
	return err
}
