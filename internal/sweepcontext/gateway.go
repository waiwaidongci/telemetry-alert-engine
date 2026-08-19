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
	<-timer.C
	return nil
}
