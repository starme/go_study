package queue

import (
	"github.com/hibiken/asynq"
	"log"
)

func Start(fn func(mux *asynq.ServeMux)) error {
	//// mux maps a type to a handler
	mux := asynq.NewServeMux()
	// ...register other handlers...
	//mux.Handle("event", Task.Handler)
	if err := server().Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
		return err
	}
	return nil
}
