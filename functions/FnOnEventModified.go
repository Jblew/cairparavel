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

// FnOnEventModified cloud function
func FnOnEventModified(ctx context.Context, e firestoreEventFnOnEventModified) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnEventModified",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventModified triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		event := e.Value.Fields.ToEvent()
		event.ID = filepath.Base(e.Value.Name)
		log.Printf("Parsed event %+v", event)
		return event.OnModified(container)
	})
}

type firestoreEventFnOnEventModified struct {
	OldValue   firestoreValueFnOnEventModified `json:"oldValue"`
	Value      firestoreValueFnOnEventModified `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// firestoreValue holds Firestore fields.
type firestoreValueFnOnEventModified struct {
	CreateTime time.Time                           `json:"createTime"`
	Fields     eventinputtypes.EventFirestoreInput `json:"fields"`
	Name       string                              `json:"name"`
	UpdateTime time.Time                           `json:"updateTime"`
}
