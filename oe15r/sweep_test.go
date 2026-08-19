package oe15r

import (
	"context"
	"errors"
	"testing"
	"time"

	sweep "github.com/example/telemetry-alert/internal/sweepcontext"
)

func TestEchoSweep015(t *testing.T) {
	t.Run("request context", TestEchoContext015)
	t.Run("cancelled retry", TestEchoRetry015)
}

func TestEchoContext015(t *testing.T) {
	service := sweep.NewService(sweep.NewRepository(sweep.NewGateway(40*time.Millisecond)), sweep.NewWorker(time.Millisecond))
	short, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	started := time.Now()
	err := service.Sweep(short)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline error, got %v", err)
	}
	if time.Since(started) > 30*time.Millisecond {
		t.Fatal("deadline was not propagated")
	}
	long, longCancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer longCancel()
	if err := service.Sweep(long); err != nil {
		t.Fatalf("later request inherited stale context: %v", err)
	}
}

func TestEchoRetry015(t *testing.T) {
	worker := sweep.NewWorker(time.Millisecond)
	service := sweep.NewService(sweep.NewRepository(sweep.NewGateway(50*time.Millisecond)), worker)
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Millisecond)
	defer cancel()
	err := service.RetrySweep(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected cancellation, got %v", err)
	}
	before := worker.Attempts()
	time.Sleep(5 * time.Millisecond)
	if worker.Attempts() != before {
		t.Fatal("worker kept retrying after cancellation")
	}
}
