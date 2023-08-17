package event

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"sync"
)

type IHandler interface {
	Broadcast(context.Context) error
	Dispatch(context.Context) error
}

type DefaultHandler struct {
	topic     string
	payload   *asynq.Task
	listeners []IListener
}

func NewHandler(topic string, payload interface{}, listeners []IListener) IHandler {
	marshal, err := json.Marshal(payload)
	if err != nil {
		return nil
	}
	task := asynq.NewTask("event", marshal)

	return &DefaultHandler{
		topic:     topic,
		payload:   task,
		listeners: listeners,
	}
}

func (d *DefaultHandler) Broadcast(ctx context.Context) error {
	errors := NewEventError()
	wg := sync.WaitGroup{}
	wg.Add(len(d.listeners))
	for _, listener := range d.listeners {
		go func(listener IListener) {
			if err := listener.ProcessTask(ctx, d.payload); err != nil {
				errors.AddError(GetStructName(listener), err)
			}
			wg.Done()
		}(listener)
	}
	wg.Wait()
	return errors.Err()
}

func (d *DefaultHandler) Dispatch(ctx context.Context) error {
	errors := NewEventError()
	for _, listener := range d.listeners {
		if err := listener.ProcessTask(ctx, d.payload); err != nil {
			errors.AddError(GetStructName(listener), err)
		}
	}
	return errors.Err()
}
