package event

import (
	"fmt"
	"strings"
	"sync"
)

type Error struct {
	mux     *sync.Mutex
	message map[string]error
}

func NewEventError() *Error {
	return &Error{
		mux:     &sync.Mutex{},
		message: make(map[string]error),
	}
}

func (e *Error) Error() string {
	var errStr []string
	for k, err := range e.message {
		errStr = append(errStr, fmt.Sprintf("[%s]%s", k, err.Error()))
	}
	if len(errStr) == 0 {
		return ""
	}
	return fmt.Sprintf("%s", strings.Join(errStr, "; "))
}

func (e *Error) AddError(key string, err error) {
	e.mux.Lock()
	defer e.mux.Unlock()

	e.message[key] = err
}

func (e *Error) HasError() bool {
	return len(e.message) > 0
}

func (e *Error) GetError(key string) error {
	return e.message[key]
}

func (e *Error) GetErrors() map[string]error {
	return e.message
}

func (e *Error) Err() error {
	if e.HasError() {
		return e
	}
	return nil
}
