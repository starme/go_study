package queue

import (
	"fmt"
	"github.com/hibiken/asynq"
	"star/listeners"
	"time"
)

func Start() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: "localhost:6379",
			DB:   0,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
			StrictPriority: true,
			RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
				return 2 * time.Second
			},
		},
	)

	mux := NewServeMux()

	mux.Handle("Event:TEvent", []asynq.Handler{listeners.TListener{}, listeners.AListener{}, listeners.VListener{}})

	if err := srv.Run(mux); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
