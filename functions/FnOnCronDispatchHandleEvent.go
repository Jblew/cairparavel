package functions

import (
	"context"
	"log"

	"github.com/Jblew/cairparavel/functions/app/apps/eventsapp"
	"github.com/Jblew/cairparavel/functions/util"
)

// FnOnCronHandleEvents cloud function
func FnOnCronHandleEvents(ctx context.Context, e PubSubMessage) {
	opts := util.FunctionHandlerOpts{
		Name:       "FnOnCronHandleEvents",
		LogErrorFn: log.Printf,
		LogPanicFn: log.Printf,
		LogDoneFn:  log.Printf,
	}
	util.FunctionHandler(opts, func() error {
		return eventsapp.OnPeriodicalHandleEvents(container)
	})
}

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
