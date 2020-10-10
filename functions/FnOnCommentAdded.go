package functions

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/app/domain"
)

// FnOnCommentAdded cloud function
func FnOnCommentAdded(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnCommentAdded triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	comment := domain.EventComment{
		ID:        e.Value.Fields.id.StringValue,
		EventID:   e.Value.Fields.eventId.StringValue,
		AuthorUID: e.Value.Fields.authorUid.StringValue,
		Contents:  e.Value.Fields.contents.StringValue,
		Time:      e.Value.Fields.time.NumberValue,
	}
	return comment.OnAdded(container)
}
