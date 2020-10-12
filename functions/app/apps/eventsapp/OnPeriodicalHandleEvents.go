package eventsapp

import (
	"log"
	"time"

	"github.com/Jblew/cairparavel/functions/app/domain"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// OnPeriodicalHandleEvents performs periodical checks of events (eg state changed based on time)
func OnPeriodicalHandleEvents(container *ioccontainer.Container) error {
	var eventsRepo domain.EventRepository
	container.Make(&eventsRepo)

	var stateRepo EventStateChangedRepository
	container.Make(&stateRepo)

	events, err := eventsRepo.GetAllNonFinishedAt(time.Now())
	if err != nil {
		return err
	}

	var lastErr error
	for _, event := range events {
		state, err := event.GetStateAt(time.Now(), container)
		if err != nil {
			log.Printf("Cannot get state of event %s: %v", event.ID, err)
			lastErr = err
			continue
		}

		previous, err := stateRepo.GetPreviousAndUpdateEventState(event.ID, state)
		if err != nil {
			log.Printf("Cannot update state of event %s: %v", event.ID, err)
			lastErr = err
			continue
		}

		if previous != state {
			err = event.OnStateChanged(previous, container)
			if err != nil {
				log.Printf("Cannot run OnStateChanged handler for event %s: %v", event.ID, err)
				lastErr = err
				continue
			}
		}
	}
	return lastErr
	return nil
}
