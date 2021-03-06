package services

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// EventObserverRepositoryFirestore implements EventObserverRepository
type EventObserverRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// GetAllForEvent retrives observers of an event
func (repo *EventObserverRepositoryFirestore) GetAllForEvent(eventID string) ([]domain.EventObserver, error) {
	if len(eventID) == 0 {
		return []domain.EventObserver{}, fmt.Errorf("InvalidArgument: Empty eventID")
	}

	collectionRef := repo.Firestore.Collection(config.FirestorePaths.ObserversForEventCol(eventID))
	snapshots, err := collectionRef.Documents(repo.Context).GetAll()
	if err != nil {
		return []domain.EventObserver{}, err
	}

	results := make([]domain.EventObserver, 0, len(snapshots))

	for _, snapshot := range snapshots {
		observer, err := observerFromSnapshot(snapshot)
		if err != nil {
			log.Printf("Invalid event observer fetched: %v", err)
		} else {
			results = append(results, observer)
		}
	}
	return results, nil
}

// DeleteAllForEvent retrives observers of an event
func (repo *EventObserverRepositoryFirestore) DeleteAllForEvent(eventID string) error {
	if len(eventID) == 0 {
		return fmt.Errorf("InvalidArgument: Empty eventID")
	}
	collectionRef := repo.Firestore.Collection(config.FirestorePaths.ObserversForEventCol(eventID))
	documentRefs, err := collectionRef.DocumentRefs(repo.Context).GetAll()
	if err != nil {
		return err
	}

	if len(documentRefs) > 0 {
		batch := repo.Firestore.Batch()
		for _, docRef := range documentRefs {
			batch = batch.Delete(docRef)
		}
		_, err = batch.Commit(repo.Context)
		return err
	}
	return nil
}

// Add saves event observer
func (repo *EventObserverRepositoryFirestore) Add(observer domain.EventObserver) error {
	if err := observer.Validate(); err != nil {
		return err
	}
	docRef := repo.Firestore.Doc(config.FirestorePaths.ObserversForEventForUserDoc(observer.EventID, observer.UID))
	_, err := docRef.Set(repo.Context, observer)
	return err
}

func observerFromSnapshot(snapshot *firestore.DocumentSnapshot) (domain.EventObserver, error) {
	var observer domain.EventObserver
	err := snapshot.DataTo(&observer)
	if err != nil {
		return domain.EventObserver{}, err
	}
	if err = observer.Validate(); err != nil {
		return domain.EventObserver{}, err
	}
	return observer, nil
}
