package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"sync"
)

type Listener interface {
	Handle(context.Context, interface{}) error
}

func event(payload interface{}) error {
	return NewBus().dispatch(NewEvent(payload))
}

func until(payload interface{}) error {
	return NewBus().dispatchSync(getStructName(payload), payload)
}

func untilEvent(event Event) error {
	return NewBus().dispatchSync(event.Name, event.Payload)
}

var (
	once = sync.Once{}
	bus  *Bus
)

type Bus struct {
	mux         *sync.Mutex
	context     context.Context
	subscribers map[string][]Listener
}

func NewBus() *Bus {
	once.Do(func() {
		bus = &Bus{
			mux:         &sync.Mutex{},
			context:     context.Background(),
			subscribers: make(map[string][]Listener),
		}
	})
	return bus
}

func (b *Bus) Subscribe(topic string, subscriber Listener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], subscriber)
}

func (b *Bus) Unsubscribe(topic string, subscriber Listener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	for i, s := range b.subscribers[topic] {
		if s == subscriber {
			b.subscribers[topic] = append(b.subscribers[topic][:i], b.subscribers[topic][i+1:]...)
			break
		}
	}
}

func (b *Bus) dispatch(event Event) error {
	marshal, err := json.Marshal(event)
	if err != nil {
		return err
	}
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
	defer func(client *asynq.Client) {
		if err := client.Close(); err != nil {

		}
	}(client)

	if _, err = client.Enqueue(asynq.NewTask("event", marshal)); err != nil {
		return err
	}
	return nil
}

func (b *Bus) dispatchSync(topic string, payload interface{}) error {
	fmt.Printf("topic: %s\nsubscribers: %#v\n", topic, b.subscribers)
	if len(b.subscribers[topic]) == 0 {
		return nil
	}

	var errors = NewEventError()

	for _, s := range b.subscribers[topic] {
		if err := s.(Listener).Handle(b.context, payload); err != nil {
			errors.AddError(getStructName(s), err)
		}
	}

	return errors.r()
}

func Handler(ctx context.Context, t *asynq.Task) error {
	var e Event
	fmt.Println("handle event task.......")
	err := json.Unmarshal(t.Payload(), &e)
	if err != nil {
		return err
	}
	fmt.Printf("event: %#v\n", e)
	return untilEvent(e)
}
