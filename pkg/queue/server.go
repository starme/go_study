package queue

import (
	"github.com/hibiken/asynq"
	"log"
)

type HandleFunc func(mux *asynq.ServeMux)

func (fn HandleFunc) Handle(mux *asynq.ServeMux) {
	fn(mux)
}

type TaskServer struct {
	mux        *asynq.ServeMux
	log        *log.Logger
	HandleFunc HandleFunc
}

func (t *TaskServer) Start() error {
	//// mux maps a type to a handler
	t.mux = asynq.NewServeMux()
	// ...register other handlers...
	//mux.Handle("event", Task.Handler)
	t.HandleFunc.Handle(t.mux)

	if err := server().Run(t.mux); err != nil {
		log.Fatalf("could not run server: %v", err)
		return err
	}
	return nil
}
