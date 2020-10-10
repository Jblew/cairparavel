package functions

import (
	"context"

	"github.com/Jblew/cairparavel/functions/app/apps/eventsapp"
)

// FnOnCronHandleEvents cloud function
func FnOnCronHandleEvents(ctx context.Context, e PubSubMessage) error {
	return eventsapp.OnPeriodicalHandleEvents(container)
}
