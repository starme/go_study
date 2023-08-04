package event

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

// Path: bus.go
func init() {
	bus = NewBus()
	bus.Subscribe("TEvent", TListener{})
	//bus.Subscribe(AListener{})
	//bus.Subscribe(VListener{})
}

type TEvent struct {
	Name string
}

type TListener struct{}

func (l TListener) Handle(_ context.Context, payload interface{}) error {
	event, ok := payload.(TEvent)
	if !ok {
		panic("payload is not TEvent")
	}

	fmt.Printf("event: %#v\n", event)
	return errors.New("event handler error")
}

func TestDispatchSync(t *testing.T) {
	event := TEvent{Name: "test"}
	err := bus.dispatchSync("TEvent", event)
	if err != nil {
		fmt.Printf("listener error: %s\n", err.Error())
	}
}

func TestEvent(t *testing.T) {
	payload := TEvent{Name: "test"}

	if err := event(payload); err != nil {
		fmt.Printf("listener error: %s\n", err.Error())
	}
}

func TestUntil(t *testing.T) {
	payload := TEvent{Name: "test"}

	if err := until(payload); err != nil {
		fmt.Printf("listener error: %s\n", err.Error())
	}
}
