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

// FnOnEventMemberSignupDeleted cloud function
func FnOnEventMemberSignupDeleted(ctx context.Context, e firestoreEventFnOnEventMemberSignupDeleted) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnEventMemberSignupDeleted",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnEventMemberSignupDeleted triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		signup := e.OldValue.Fields.ToEventSignup()
		signup.UID = filepath.Base(e.OldValue.Name)
		log.Printf("Parsed signup %+v", signup)
		return signup.OnDeleted(container)
	})
}

type firestoreEventFnOnEventMemberSignupDeleted struct {
	OldValue   firestoreValueFnOnEventMemberSignupDeleted `json:"oldValue"`
	Value      firestoreValueFnOnEventMemberSignupDeleted `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type firestoreValueFnOnEventMemberSignupDeleted struct {
	CreateTime time.Time                                 `json:"createTime"`
	Fields     eventinputtypes.EventSignupFirestoreInput `json:"fields"`
	Name       string                                    `json:"name"`
	UpdateTime time.Time                                 `json:"updateTime"`
}
