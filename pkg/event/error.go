package event

import (
	"fmt"
	"strings"
	"sync"
)

type EventError struct {
	mux     *sync.Mutex
	message map[string]error
}

func NewEventError() *EventError {
	return &EventError{
		mux:     &sync.Mutex{},
		message: make(map[string]error),
	}
}

func (e *EventError) Error() string {
	var errStr []string
	for k, err := range e.message {
		errStr = append(errStr, fmt.Sprintf("[%s]%s", k, err.Error()))
	}
	if len(errStr) == 0 {
		return ""
	}
	return fmt.Sprintf("%s", strings.Join(errStr, "; "))
}

func (e *EventError) AddError(key string, err error) {
	e.mux.Lock()
	defer e.mux.Unlock()

	e.message[key] = err
}

func (e *EventError) HasError() bool {
	return len(e.message) > 0
}

func (e *EventError) GetError(key string) error {
	return e.message[key]
}

func (e *EventError) GetErrors() map[string]error {
	return e.message
}

func (e *EventError) r() error {
	if e.HasError() {
		return e
	}
	return nil
}
