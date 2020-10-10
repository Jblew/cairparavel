package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/eventinputtypes"
)

// FnOnEventMemberSignupCreated cloud function
func FnOnEventMemberSignupCreated(ctx context.Context, e firestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnEventMemberSignupCreated triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	signup := e.Value.Fields.ToEventSignup()
	return signup.OnAdded(container)
}

type firestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type firestoreValue struct {
	CreateTime time.Time                                 `json:"createTime"`
	Fields     eventinputtypes.EventSignupFirestoreInput `json:"fields"`
	Name       string                                    `json:"name"`
	UpdateTime time.Time                                 `json:"updateTime"`
}
