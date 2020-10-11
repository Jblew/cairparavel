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

// FnOnCommentAdded cloud function
func FnOnCommentAdded(ctx context.Context, e firestoreEventFnOnCommentAdded) error {
	opts := util.FunctionHandlerOpts{
		Name:       "FnOnCommentAdded",
		LogErrorFn: log.Printf,
		LogPanicFn: log.Printf,
		LogDoneFn:  log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnCommentAdded triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		comment := e.Value.Fields.ToEventComment()
		log.Printf("Parsed comment %+v", comment)
		return comment.OnAdded(container)
	})
}

type firestoreEventFnOnCommentAdded struct {
	OldValue   firestoreValueFnOnCommentAdded `json:"oldValue"`
	Value      firestoreValueFnOnCommentAdded `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type firestoreValueFnOnCommentAdded struct {
	CreateTime time.Time                                  `json:"createTime"`
	Fields     eventinputtypes.EventCommentFirestoreInput `json:"fields"`
	Name       string                                     `json:"name"`
	UpdateTime time.Time                                  `json:"updateTime"`
}
