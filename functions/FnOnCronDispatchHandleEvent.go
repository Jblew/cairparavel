package functions

import (
	"context"

	"github.com/Jblew/cairparavel/functions/app/apps/eventsapp"
)

// FnOnCronHandleEvents cloud function
func FnOnCronHandleEvents(ctx context.Context, e PubSubMessage) error {
	return eventsapp.OnPeriodicalHandleEvents(container)
}

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
