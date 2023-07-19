package bus

import (
	"context"
)

type IListener interface {
	Handler(ctx context.Context, task interface{}) error
}
