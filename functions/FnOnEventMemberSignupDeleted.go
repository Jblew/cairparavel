package functions

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// FnOnEventMemberSignupDeleted cloud function
func FnOnEventMemberSignupDeleted(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnEventMemberSignupDeleted triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	signup := domain.EventSignup{
		UID:         e.Value.Fields.uid.StringValue,
		EventID:     e.Value.Fields.eventId.StringValue,
		DisplayName: e.Value.Fields.displayName.StringValue,
	}
	return signup.OnDeleted(container)
}
