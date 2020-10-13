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

// FnOnEventVoteDeleted cloud function
func FnOnEventVoteDeleted(ctx context.Context, e firestoreEventFnOnEventVoteDeleted) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnEventVoteDeleted",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventVoteDeleted triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		votes := e.OldValue.Fields.ToEventTimeVotes()
		votes.UID = filepath.Base(e.OldValue.Name)
		log.Printf("Parsed votes %+v", votes)
		return votes.OnDeleted(container)
	})
}

type firestoreEventFnOnEventVoteDeleted struct {
	OldValue   firestoreValueFnOnEventVoteDeleted `json:"oldValue"`
	Value      firestoreValueFnOnEventVoteDeleted `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// firestoreValue holds Firestore fields.
type firestoreValueFnOnEventVoteDeleted struct {
	CreateTime time.Time                                    `json:"createTime"`
	Fields     eventinputtypes.EventTimeVotesFirestoreInput `json:"fields"`
	Name       string                                       `json:"name"`
	UpdateTime time.Time                                    `json:"updateTime"`
}
