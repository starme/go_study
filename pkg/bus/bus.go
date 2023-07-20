package bus

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"sync"
)

var bus *Bus

// Bus is an interface for event listeners.
type Bus struct {
	mux         *sync.Mutex
	client      *asynq.Client
	subscribers map[string][]IListener
}

// NewBus returns a new event bus.
func NewBus() *Bus {
	return &Bus{
		mux:         &sync.Mutex{},
		client:      asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"}),
		subscribers: make(map[string][]IListener),
	}
}

// Subscribe adds a new event listener.
func (b *Bus) Subscribe(topic string, subscriber IListener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], subscriber)
}

// Unsubscribe removes an event listener.
func (b *Bus) Unsubscribe(topic string, subscriber IListener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	for i, s := range b.subscribers[topic] {
		if s == subscriber {
			b.subscribers[topic] = append(b.subscribers[topic][:i], b.subscribers[topic][i+1:]...)
			break
		}
	}
}

// Publish publishes an event to all subscribers.
func (b *Bus) Publish(topic string, payload interface{}, opts ...Option) error {
	b.mux.Lock()
	defer b.mux.Unlock()

	opt, err := composeOptions(opts...)
	if err != nil {
		return err
	}

	if opt.exec == "serial" {
		return b.publishSerial(topic, payload)
	}

	return b.publishParallel(topic, payload)
}

func (b *Bus) publishSerial(topic string, payload interface{}) error {
	for _, subscriber := range b.subscribers[topic] {
		fmt.Printf("Publishing event to subscriber. %#v, %#v\n", subscriber, payload)
		marshal, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		return subscriber.Handler(context.TODO(), marshal)
	}
	return nil
}

func (b *Bus) publishParallel(topic string, payload interface{}) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(b.subscribers[topic]))
	resCh := make(chan error, len(b.subscribers[topic]))
	for _, subscriber := range b.subscribers[topic] {
		go func(subscriber IListener, resCh chan<- error) {
			defer wg.Done()
			fmt.Printf("Publishing event to subscriber. %#v\n", subscriber)
			//subscriber.Handler(payload)
			marshal, err := json.Marshal(payload)
			if err != nil {
				resCh <- err
				return
			}
			resCh <- subscriber.Handler(context.TODO(), marshal)
		}(subscriber, resCh)
	}
	wg.Wait()
	close(resCh)
	for err := range resCh {
		if err != nil {
			return err
		}
	}
	return nil
}
