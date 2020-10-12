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

// FnOnNotificationQueued cloud function
func FnOnNotificationQueued(ctx context.Context, e firestoreEventFnOnNotificationQueued) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnNotificationQueued",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		meta, err := metadata.FromContext(ctx)
		if err != nil {
			return fmt.Errorf("metadata.FromContext: %v", err)
		}
		log.Printf("Function FnOnNotificationQueued triggered by change to: %v", meta.Resource)
		log.Printf("Old value: %+v", e.OldValue)
		log.Printf("New value: %+v", e.Value)

		notification, err := e.Value.Fields.ToNotification()
		notification.ID = filepath.Base(e.Value.Name)
		if err != nil {
			return err
		}
		log.Printf("Parsed notification %+v", notification)
		return notification.OnQueued(container)
	})
}

type firestoreEventFnOnNotificationQueued struct {
	OldValue   firestoreValueFnOnNotificationQueued `json:"oldValue"`
	Value      firestoreValueFnOnNotificationQueued `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type firestoreValueFnOnNotificationQueued struct {
	CreateTime time.Time                                  `json:"createTime"`
	Fields     eventinputtypes.NotificationFirestoreInput `json:"fields"`
	Name       string                                     `json:"name"`
	UpdateTime time.Time                                  `json:"updateTime"`
}
