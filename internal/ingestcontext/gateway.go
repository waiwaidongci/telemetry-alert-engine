package ingestcontext

import (
	"context"
	"errors"
)

var ErrGatewayUnavailable = errors.New("ingest gateway unavailable")

type Gateway struct{}

func (Gateway) Load(ctx context.Context, ready <-chan struct{}) error {
	select {
	case <-ready:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
