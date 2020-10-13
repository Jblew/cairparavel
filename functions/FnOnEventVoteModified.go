package functions

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/Jblew/cairparavel/functions/eventinputtypes"
	"github.com/Jblew/cairparavel/functions/util"
)

// FnOnEventVoteModified cloud function
func FnOnEventVoteModified(ctx context.Context, e firestoreEventFnOnEventVoteModified) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnEventVoteModified",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventVoteModified triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		votes := e.Value.Fields.ToEventTimeVotes()
		votes.UID = filepath.Base(e.Value.Name)
		log.Printf("Parsed votes %+v", votes)
		return votes.OnModified(container)
	})
}

type firestoreEventFnOnEventVoteModified struct {
	OldValue   firestoreValueFnOnEventVoteModified `json:"oldValue"`
	Value      firestoreValueFnOnEventVoteModified `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// firestoreValue holds Firestore fields.
type firestoreValueFnOnEventVoteModified struct {
	CreateTime time.Time                                    `json:"createTime"`
	Fields     eventinputtypes.EventTimeVotesFirestoreInput `json:"fields"`
	Name       string                                       `json:"name"`
	UpdateTime time.Time                                    `json:"updateTime"`
}
