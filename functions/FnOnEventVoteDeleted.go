package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/eventinputtypes"
)

// FnOnEventVoteDeleted cloud function
func FnOnEventVoteDeleted(ctx context.Context, e firestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function FnOnEventVoteDeleted triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)

	votes := e.Value.Fields.ToEventTimeVotes()
	log.Printf("Parsed votes %+v", votes)
	return votes.OnCreated(container)
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
	CreateTime time.Time                                    `json:"createTime"`
	Fields     eventinputtypes.EventTimeVotesFirestoreInput `json:"fields"`
	Name       string                                       `json:"name"`
	UpdateTime time.Time                                    `json:"updateTime"`
}
