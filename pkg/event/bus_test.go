package event

import (
	"fmt"
	"star/listeners"
	"testing"
)

var bus Bus

// Path: bus.go
func init() {
	bus.Register(map[string][]IListener{
		"TEvent": []IListener{
			listeners.TListener{},
			listeners.AListener{},
			listeners.VListener{},
		},
	})
}

func TestBus_Register(t *testing.T) {
	bus.Register(map[string][]IListener{
		"TEvent": {
			listeners.TListener{},
			listeners.AListener{},
			listeners.VListener{},
		},
	})

	fmt.Printf("bus: %#v\n", bus)
}

func TestBus_Broadcast(t *testing.T) {
	event := listeners.TEvent{
		Name: "TEvent",
	}
	err := bus.broadcast("TEvent", event)
	fmt.Printf("err: %v\n", err)
}

func TestBus_Dispatch(t *testing.T) {
	event := listeners.TEvent{
		Name: "TEvent",
	}
	err := bus.dispatch("TEvent", event)
	fmt.Printf("err: %v\n", err)
}
