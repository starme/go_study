package queue

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"star/listeners"
	"testing"
)

func TestNewServeMux(t *testing.T) {
	client := asynq.NewClient(&asynq.RedisClientOpt{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer func(client *asynq.Client) {
		if err := client.Close(); err != nil {
			fmt.Printf("could not enqueue task: %v", err)
		}
	}(client)
	event := listeners.TEvent{
		Name: "TEvent",
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return
	}

	task := asynq.NewTask("TEvent", payload)
	// Process the task immediately.
	info, err := client.Enqueue(task, asynq.MaxRetry(3), asynq.Queue("critical"))
	if err != nil {
		fmt.Printf("could not enqueue task: %v", err)
	}
	fmt.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
