package event

import (
	"context"
	"github.com/hibiken/asynq"
	"sync"
)

type IListener asynq.Handler

type Bus struct {
	mux         *sync.RWMutex
	context     context.Context
	subscribers map[string][]IListener
}

func (b *Bus) Register(subscribers map[string][]IListener) *Bus {
	if b.isNull() {
		b.boot()
	}

	b.subscribers = subscribers

	return b
}

func (b *Bus) boot() {
	b.mux = &sync.RWMutex{}
	b.context = context.Background()
	b.subscribers = make(map[string][]IListener)
}

func (b *Bus) subscribe(topic string, subscriber IListener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], subscriber)
}

func (b *Bus) unsubscribe(topic string, subscriber IListener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if _, ok := b.show(topic); !ok {
		return
	}

	for i, s := range b.subscribers[topic] {
		if s == subscriber {
			b.subscribers[topic] = append(b.subscribers[topic][:i], b.subscribers[topic][i+1:]...)
			break
		}
	}
}

func (b *Bus) show(topic string) ([]IListener, bool) {
	b.mux.RLock()
	defer b.mux.RUnlock()

	if _, ok := b.subscribers[topic]; !ok {
		return []IListener{}, false
	}

	return b.subscribers[topic], true
}

func (b *Bus) isNull() bool {
	return len(b.subscribers) == 0
}

func (b *Bus) broadcast(topic string, payload interface{}) error {
	handler := NewHandler(topic, payload, b.subscribers[topic])

	return handler.Broadcast(b.context)
}

func (b *Bus) dispatch(topic string, payload interface{}) error {
	handler := NewHandler(topic, payload, b.subscribers[topic])

	return handler.Dispatch(b.context)
}
