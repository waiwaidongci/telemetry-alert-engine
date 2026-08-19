package dispatchcontext

import (
	"context"
	"errors"
	"time"
)

var ErrDispatchDeadline = errors.New("dispatch deadline reached")

func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeoutCause(parent, timeout, ErrDispatchDeadline)
}
