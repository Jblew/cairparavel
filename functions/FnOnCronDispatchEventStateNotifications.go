package functions

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// FnOnCronDispatchEventStateNotifications cloud function
func FnOnCronDispatchEventStateNotifications(ctx context.Context, e PubSubMessage) error {
	log.Printf("Function FnOnCronDispatchEventStateNotifications triggered")
	return nil
}
