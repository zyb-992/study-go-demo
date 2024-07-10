package observer

import (
	"context"
	"fmt"
)

type SyncEventBus struct {
	baseEventBus *BaseEventBus
}

func NewSyncEventbus() *SyncEventBus {
	return &SyncEventBus{
		baseEventBus: NewEventBus(),
	}
}

func (s *SyncEventBus) Publish(ctx context.Context, event *Event) {
	subscribers, ok := s.baseEventBus.observers[event.Topic]
	if !ok {
		s.baseEventBus.observers[event.Topic] = make(map[Observer]struct{})
		return
	}

	handleErrs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.OnChange(ctx, event); err != nil {
			handleErrs[subscriber] = err
		}
	}

	s.handleErrors(handleErrs)
}

func (s *SyncEventBus) handleErrors(errs map[Observer]error) {
	for o, err := range errs {
		fmt.Printf("Publish to subscriber [%s] err [%v]", o.Name(), err.Error())
	}

}
