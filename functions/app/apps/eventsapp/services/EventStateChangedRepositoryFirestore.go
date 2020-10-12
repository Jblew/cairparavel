package services

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Jblew/cairparavel/functions/app/config"
	"github.com/Jblew/cairparavel/functions/app/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EventStateChangedRepositoryFirestore implements EventStateChangedRepository
type EventStateChangedRepositoryFirestore struct {
	Firestore *firestore.Client
	Context   context.Context
}

// GetPreviousAndUpdateEventState retrives count of signups
func (repo *EventStateChangedRepositoryFirestore) GetPreviousAndUpdateEventState(eventID string, state domain.EventState) (domain.EventState, error) {
	docPath := config.FirestorePaths.EventLastStateDoc(eventID)
	docRef := repo.Firestore.Doc(docPath)

	snapshot, err := docRef.Get(repo.Context)
	if err != nil && status.Code(err) != codes.NotFound {
		return domain.EventStateNonexistent, err
	}
	previousState := domain.EventStateNonexistent

	if snapshot.Exists() {
		var result stateDoc
		err = snapshot.DataTo(&result)
		if err != nil {
			return previousState, err
		}
		previousState = result.State
	}

	newStateDoc := stateDoc{
		State: state,
	}
	_, err = repo.Firestore.Doc(docPath).Set(repo.Context, newStateDoc)
	if err != nil {
		return previousState, err
	}
	return previousState, nil
}

type stateDoc struct {
	State domain.EventState `json:"state"`
}
