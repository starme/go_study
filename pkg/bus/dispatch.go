package bus

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

type OptionType int

const (
	ExecOpt OptionType = iota
)

const (
	// Default exec type is serial.
	defaultExecType = "serial"
)

type Option interface {
	// String returns a string representation of the option.
	String() string

	// Type describes the type of the option.
	Type() OptionType

	// Value returns a value used to create this option.
	Value() interface{}
}

type (
	execOption string
)

// SerialExec is an option to execute event handler synchronously.
func SerialExec() Option {
	return execOption("serial")
}

// ParallelExec AsyncExec is an option to execute event handler asynchronously.
func ParallelExec() Option {
	return execOption("parallel")
}

func (e execOption) String() string     { return fmt.Sprintf("execOption(%s)", e) }
func (e execOption) Type() OptionType   { return ExecOpt }
func (e execOption) Value() interface{} { return string(e) }

type option struct {
	exec string
}

// composeOptions merges user provided options into the default options
// and returns the composed option.
// It also validates the user provided options and returns an error if any of
// the user provided options fail the validations.
func composeOptions(opts ...Option) (option, error) {
	res := option{
		exec: defaultExecType,
	}
	for _, opt := range opts {
		switch opt := opt.(type) {
		case execOption:
			res.exec = opt.Value().(string)
		default:
			// ignore unexpected option
		}
	}
	return res, nil
}

func DispatchSync(event IEvent, ops ...Option) error {
	if bus == nil {
		bus = NewBus()
	}

	return bus.Publish(event.Name(), event, ops...)
}

func Dispatch(event IEvent, opts ...Option) {
	if bus == nil {
		bus = NewBus()
	}

	task, err := json.Marshal(event)
	if err != nil {
		panic("Stoxx")
	}

	info, err := bus.client.Enqueue(asynq.NewTask(event.Name(), task))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	//return bus.Publish(event.Name(), event, opts...)
}

func Register(event IEvent, subscriber IListener) {
	if bus == nil {
		bus = NewBus()
	}

	bus.Subscribe(event.Name(), subscriber)
}
