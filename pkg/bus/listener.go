package bus

import (
	"context"
)

type IListener interface {
	Handler(ctx context.Context, payload []byte) error
}
