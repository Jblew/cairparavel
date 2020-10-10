package functions

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// FnOnEventVoteCreated cloud function
func FnOnEventVoteCreated(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnEventVoteCreated triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	times := make([]int64)

	for _, firestoreValue := range e.Value.Fields.times.ArrayValue.values {
		append(times, firestoreValue.NumberValue)
	}

	votes := domain.EventTimeVotes{
		UID:         e.Value.Fields.uid.StringValue,
		EventID:     e.Value.Fields.eventId.StringValue,
		DisplayName: e.Value.Fields.displayName.StringValue,
		Times:       times,
	}
	return votes.OnCreated(container)
}
