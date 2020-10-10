package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/eventinputtypes"
)

// FnOnEventModified cloud function
func FnOnEventModified(ctx context.Context, e firestoreEvent) error {
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

	event := e.Value.Fields.ToEvent()
	return event.OnModified(container)
}

type firestoreEvent struct {
	OldValue   firestoreValue `json:"oldValue"`
	Value      firestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// firestoreValue holds Firestore fields.
type firestoreValue struct {
	CreateTime time.Time                           `json:"createTime"`
	Fields     eventinputtypes.EventFirestoreInput `json:"fields"`
	Name       string                              `json:"name"`
	UpdateTime time.Time                           `json:"updateTime"`
}
