package ingestcontext

import (
	"context"
	"errors"
)

var ErrGatewayUnavailable = errors.New("ingest gateway unavailable")

type Gateway struct{}

func (Gateway) Load(ctx context.Context, ready <-chan struct{}) error {
	_ = ctx
	select {
	case <-ready:
		return nil
	default:
		return nil
	}
}
