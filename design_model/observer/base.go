package observer

import (
	"context"
	"fmt"
	"sync"
)

// Base Definition

// BaseObserver 基类观察者
type BaseObserver struct {
	name string
}

func NewBaseObserver(name string) *BaseObserver {
	o := &BaseObserver{
		name: name,
	}

	return o
}

func (o *BaseObserver) OnChange(ctx context.Context, event *Event) error {
	fmt.Printf("[%s]: listen change on event [%s]", o.name, event.Topic)
	return nil
}

func (o *BaseObserver) Name() string {
	return o.name
}

type BaseEventBus struct {
	mutex     sync.RWMutex
	observers map[string]map[Observer]struct{}
}

func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if _, ok := b.observers[topic]; !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}

	b.observers[topic][o] = struct{}{}
}

func (b *BaseEventBus) UnSubscribe(topic string, o Observer) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	delete(b.observers[topic], o)
}

func NewEventBus() *BaseEventBus {
	b := &BaseEventBus{
		mutex:     sync.RWMutex{},
		observers: make(map[string]map[Observer]struct{}),
	}

	return b
}
