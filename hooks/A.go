package hooks

import (
	"context"
	"github.com/hibiken/asynq"
	"star/pkg/log"
	"star/pkg/queue"
)

type AHook struct {
	Name string
	Age  int
}

var ATask = &queue.Task{
	Name: "A",
	Payload: &AHook{
		Name: "aaaa",
		Age:  20,
	},
	Handler: func(ctx context.Context, task *asynq.Task) error {
		log.Info("xxxxxxxxxxxxaaaaaaaaaa")
		return nil
	},
}

var ATaskOptions = []asynq.Option{
	asynq.MaxRetry(3),
}

func (h *AHook) GetName() string {
	return h.Name
}

func (h *AHook) GetAge() int {
	return h.Age
}

func (h *AHook) Middleware() []string {
	return []string{"A", "B"}
}
