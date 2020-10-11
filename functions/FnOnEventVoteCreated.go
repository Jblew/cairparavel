package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/eventinputtypes"
	"github.com/Jblew/cairparavel/functions/util"
)

// FnOnEventVoteCreated cloud function
func FnOnEventVoteCreated(ctx context.Context, e firestoreEventFnOnEventVoteCreated) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnEventVoteCreated",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventVoteCreated triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		votes := e.Value.Fields.ToEventTimeVotes()
		log.Printf("Parsed votes %+v", votes)
		return votes.OnAdded(container)
	})
}

type firestoreEventFnOnEventVoteCreated struct {
	OldValue   firestoreValueFnOnEventVoteCreated `json:"oldValue"`
	Value      firestoreValueFnOnEventVoteCreated `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// firestoreValue holds Firestore fields.
type firestoreValueFnOnEventVoteCreated struct {
	CreateTime time.Time                                    `json:"createTime"`
	Fields     eventinputtypes.EventTimeVotesFirestoreInput `json:"fields"`
	Name       string                                       `json:"name"`
	UpdateTime time.Time                                    `json:"updateTime"`
}
