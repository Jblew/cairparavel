package functions

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
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
	return nil
}
