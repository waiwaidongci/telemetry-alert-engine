package sweepcontext

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestSweepHonorsEachRequestContext(t *testing.T) {
	service := NewService(NewRepository(NewGateway(40*time.Millisecond)), NewWorker(time.Millisecond))
	short, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	started := time.Now()
	err := service.Sweep(short)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline error, got %v", err)
	}
	if time.Since(started) > 30*time.Millisecond {
		t.Fatal("deadline was not propagated to gateway")
	}
	long, longCancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer longCancel()
	if err := service.Sweep(long); err != nil {
		t.Fatalf("later request inherited stale context: %v", err)
	}
}

func TestRetryStopsAfterCancellation(t *testing.T) {
	worker := NewWorker(time.Millisecond)
	service := NewService(NewRepository(NewGateway(50*time.Millisecond)), worker)
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
