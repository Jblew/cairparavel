package functions

import (
	"context"
	"log"

	"github.com/Jblew/cairparavel/functions/app/apps/eventsapp"
	"github.com/Jblew/cairparavel/functions/util"
)

// FnOnCronHandleEvents cloud function
func FnOnCronHandleEvents(ctx context.Context, e PubSubMessage) error {
	opts := util.FunctionHandlerOpts{
		Name:   "FnOnCronHandleEvents",
		Log:    log.Printf,
		LogErr: log.Printf,
	}
	return util.FunctionHandler(opts, func() error {
		return eventsapp.OnPeriodicalHandleEvents(container)
	})
}

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
