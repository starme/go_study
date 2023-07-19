package providers

import (
	"star/hooks"
	"star/pkg/bus"
)

var event = map[bus.IEvent][]bus.IListener{
	hooks.TEvent{}: {
		hooks.TListener{},
		hooks.SListener{},
		hooks.VListener{},
		hooks.AListener{},
	},
}

func Register() {
	for topic, listeners := range event {
		for _, listener := range listeners {
			bus.Register(topic, listener)
		}
	}
}

func Listeners() map[bus.IEvent][]bus.IListener {
	return event
}
