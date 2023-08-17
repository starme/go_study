package listeners

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
)

type TEvent struct {
	Name string `json:"name"`
}

type AListener struct{}

func (a AListener) ProcessTask(_ context.Context, task *asynq.Task) error {
	var event TEvent
	if err := json.Unmarshal(task.Payload(), &event); err != nil {
		return err
	}

	fmt.Printf("listener: %#v\n", a)
	return nil
}

type VListener struct{}

func (v VListener) ProcessTask(_ context.Context, task *asynq.Task) error {
	var event TEvent
	if err := json.Unmarshal(task.Payload(), &event); err != nil {
		return err
	}

	fmt.Printf("listener: %#v\n", v)
	return errors.New("a listener error: VListener")
}

type TListener struct{}

func (l TListener) ProcessTask(_ context.Context, task *asynq.Task) error {
	var event TEvent
	if err := json.Unmarshal(task.Payload(), &event); err != nil {
		return err
	}

	fmt.Printf("listener: %#v\n", l)
	return nil
}
