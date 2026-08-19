package sweepcontext

import (
	"context"
	"time"
)

type Gateway struct {
	delay time.Duration
}

func NewGateway(delay time.Duration) *Gateway {
	return &Gateway{delay: delay}
}

func (g *Gateway) Fetch(ctx context.Context) error {
	timer := time.NewTimer(g.delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
