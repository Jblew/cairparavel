package functions

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// FnOnEventModified cloud function
func FnOnEventModified(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnEventModified triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	eventID := e.Value.Fields.id.StringValue
	if len(eventID) == 0 {
		return fmt.Errorf("Empty event ID")
	}

	var eventRepo domain.EventRepository
	container.Make(&eventRepo)

	event, err := eventRepo.GetEventByID(eventID)
	if err != nil {
		return err
	}

	return event.OnModified(container)
}
