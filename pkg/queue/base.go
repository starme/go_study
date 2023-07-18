package queue

import "github.com/hibiken/asynq"

func broker() *asynq.RedisClientOpt {
	return &asynq.RedisClientOpt{
		Network:      "",
		Addr:         "localhost:6379",
		Username:     "",
		Password:     "",
		DB:           1,
		DialTimeout:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		PoolSize:     0,
		TLSConfig:    nil,
	}
}

func server() *asynq.Server {
	sev := asynq.NewServer(
		broker(),
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
		})
	return sev
}

func client() (*asynq.Client, func()) {
	c := asynq.NewClient(broker())
	cancel := func() {
		err := c.Close()
		if err != nil {

		}
	}
	return c, cancel
}
