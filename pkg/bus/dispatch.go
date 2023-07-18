package bus

import "sync"

var bus *Bus

type Bus struct {
	mux         *sync.Mutex
	subscribers map[string][]IListener
}

func NewBus() *Bus {
	return &Bus{
		mux:         &sync.Mutex{},
		subscribers: make(map[string][]IListener),
	}
}

func (b *Bus) Subscribe(topic string, subscriber IListener) {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], subscriber)
}

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

func (b *Bus) Publish(topic string, message string) {
	b.mux.Lock()
	defer b.mux.Unlock()

	for _, s := range b.subscribers[topic] {
		go func(s IListener) {
			s.Handler(message)
		}(s)
	}
}

func Dispatch(topic string, message string) {
	if bus == nil {
		bus = NewBus()
	}
	bus.Publish(topic, message)
}

func Register(topic string, subscriber IListener) {
	if bus == nil {
		bus = NewBus()
	}
	bus.Subscribe(topic, subscriber)
}
