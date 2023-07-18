package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"log"
)

func Dispatch(hook *Task, opts ...asynq.Option) {
	q, cancel := client()
	defer cancel()

	hook.Id = uuid.NewString()

	task, err := json.Marshal(hook)
	if err != nil {
		panic("Stoxx")
	}

	info, err := q.Enqueue(asynq.NewTask("event", task), opts...)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func HandleHook(ctx context.Context, task *asynq.Task) error {
	var job Task
	err := json.Unmarshal(task.Payload(), &job)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(job)
	return nil
}

type Task struct {
	Id      string
	Payload interface{}
}

func (t Task) UnmarshalJSON(data []byte) error {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	t.Id = result["id"].(string)
	t.Payload = result["payload"]
	return nil
}

func (t Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":      t.Id,
		"payload": "t.Payload",
	})
}

func client() (client *asynq.Client, cancel func()) {
	client = asynq.NewClient(&asynq.RedisClientOpt{
		Addr: "localhost:6379",
		DB:   1,
	})
	cancel = func() {
		err := client.Close()
		if err != nil {

		}
	}
	return client, cancel
}
