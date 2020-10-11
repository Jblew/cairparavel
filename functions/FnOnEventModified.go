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

// FnOnEventModified cloud function
func FnOnEventModified(ctx context.Context, e firestoreEventFnOnEventModified) {
	opts := util.FunctionHandlerOpts{
		Name:       "FnOnEventModified",
		LogErrorFn: log.Printf,
		LogPanicFn: log.Printf,
		LogDoneFn:  log.Printf,
	}
	util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventModified triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		event := e.Value.Fields.ToEvent()
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
