package queue

import (
	"github.com/hibiken/asynq"
	"log"
	"star/pkg/event"
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
	t.mux.HandleFunc("event", event.Handler)
	//t.HandleFunc.Handle(t.mux)

	if err := server().Run(t.mux); err != nil {
		log.Fatalf("could not run server: %v", err)
		return err
	}
	return nil
}
