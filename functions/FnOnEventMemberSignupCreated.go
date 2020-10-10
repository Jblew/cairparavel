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
func FnOnEventMemberSignupCreated(ctx context.Context, e firestoreEventFnOnEventMemberSignupCreated) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnEventMemberSignupCreated triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	signup := e.Value.Fields.ToEventSignup()
	log.Printf("Parsed signup %+v", signup)
	return signup.OnAdded(container)
}

type firestoreEventFnOnEventMemberSignupCreated struct {
	OldValue   firestoreValueFnOnEventMemberSignupCreated `json:"oldValue"`
	Value      firestoreValueFnOnEventMemberSignupCreated `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type firestoreValueFnOnEventMemberSignupCreated struct {
	CreateTime time.Time                                 `json:"createTime"`
	Fields     eventinputtypes.EventSignupFirestoreInput `json:"fields"`
	Name       string                                    `json:"name"`
	UpdateTime time.Time                                 `json:"updateTime"`
}
