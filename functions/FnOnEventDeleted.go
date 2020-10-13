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

// FnOnEventDeleted cloud function
func FnOnEventDeleted(ctx context.Context, e firestoreEventFnOnEventDeleted) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnEventDeleted",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventDeleted triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		event := e.OldValue.Fields.ToEvent()
		event.ID = filepath.Base(e.Value.Name)
		log.Printf("Parsed event %+v", event)
		return event.OnDeleted(container)
	})
}

type firestoreEventFnOnEventDeleted struct {
	OldValue   firestoreValueFnOnEventDeleted `json:"oldValue"`
	Value      firestoreValueFnOnEventDeleted `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// firestoreValue holds Firestore fields.
type firestoreValueFnOnEventDeleted struct {
	CreateTime time.Time                           `json:"createTime"`
	Fields     eventinputtypes.EventFirestoreInput `json:"fields"`
	Name       string                              `json:"name"`
	UpdateTime time.Time                           `json:"updateTime"`
}
