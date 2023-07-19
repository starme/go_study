package hooks

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"star/pkg/log"
)

type TEvent struct{}

func (e TEvent) Name() string {
	return "TEvent"
}

type TListener struct{}

func (l TListener) Handler(_ context.Context, event interface{}) error {
	fmt.Printf("event: %#v\n", event)
	log.Info("Handler: ", zap.Any("event", event))
	return nil
}

type AListener struct{}

func (s AListener) Handler(_ context.Context, event interface{}) error {
	fmt.Printf("event: %#v\n", event)
	log.Info("Handler: ", zap.Any("event", event))
	return nil
}

type VListener struct{}

func (s VListener) Handler(_ context.Context, event interface{}) error {
	fmt.Printf("event: %#v\n", event)
	log.Info("Handler: ", zap.Any("event", event))
	return nil
}

type SListener struct{}

func (s SListener) Handler(_ context.Context, event interface{}) error {
	fmt.Printf("event: %#v\n", event)
	log.Info("Handler: ", zap.Any("event", event))
	return nil
}
