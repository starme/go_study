package queue

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"log"
)

func Dispatch(hook *Task, opts ...asynq.Option) {
	q, cancel := client()
	defer cancel()

	hook.id = uuid.NewString()

	task, err := json.Marshal(hook)
	if err != nil {
		panic("Stoxx")
	}
	
	info, err := q.Enqueue(asynq.NewTask(hook.Name, task), opts...)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}

	//mux.HandleFunc(hook.Name, hook.Handler)
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

type Task struct {
	id      string
	Name    string
	Payload interface{}
	Handler func(context.Context, *asynq.Task) error
}

func (t Task) UnmarshalJSON(data []byte) error {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	t.id = result["id"].(string)
	t.Name = result["name"].(string)
	t.Payload = result["payload"]
	return nil
}

func (t Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":      t.id,
		"name":    t.Name,
		"payload": t.Payload,
	})
}
