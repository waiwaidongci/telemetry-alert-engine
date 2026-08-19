package dispatchcontext

import (
	"context"
	"sync"
	"time"
)

type Gateway struct {
	delay time.Duration
	mu    sync.Mutex
	sent  []string
}

func NewGateway(delay time.Duration) *Gateway {
	return &Gateway{delay: delay}
}

func (g *Gateway) SetDelay(delay time.Duration) {
	g.mu.Lock()
	g.delay = delay
	g.mu.Unlock()
}

func (g *Gateway) Send(ctx context.Context, target string) error {
	g.mu.Lock()
	delay := g.delay
	g.mu.Unlock()
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	case <-timer.C:
		g.mu.Lock()
		g.sent = append(g.sent, target)
		g.mu.Unlock()
		return nil
	}
}

func (g *Gateway) Sent() []string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return append([]string(nil), g.sent...)
}
