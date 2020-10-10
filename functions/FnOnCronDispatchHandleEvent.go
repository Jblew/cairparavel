package functions

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// FnOnCronDispatchHandleEvent cloud function
func FnOnCronDispatchHandleEvent(ctx context.Context, e PubSubMessage) error {
	log.Printf("Function FnOnCronDispatchHandleEvent triggered")
	return nil
}
