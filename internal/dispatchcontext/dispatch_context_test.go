package dispatchcontext

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDispatchPropagatesTimeoutCauseToFanout(t *testing.T) {
	gateway := NewGateway(40 * time.Millisecond)
	service := NewService(NewWorker(gateway), 5*time.Millisecond)
	started := time.Now()
	err := service.Dispatch(context.Background(), []string{"pager", "webhook", "email"})
	if !errors.Is(err, ErrDispatchDeadline) {
		t.Fatalf("timeout cause was lost: %v", err)
	}
	if time.Since(started) > 30*time.Millisecond {
		t.Fatal("fanout ignored dispatch timeout")
	}
	time.Sleep(45 * time.Millisecond)
	if got := gateway.Sent(); len(got) != 0 {
		t.Fatalf("cancelled deliveries continued: %#v", got)
	}
}

func TestLaterDispatchDoesNotInheritCancellation(t *testing.T) {
	gateway := NewGateway(time.Millisecond)
	service := NewService(NewWorker(gateway), 50*time.Millisecond)
	if err := service.Dispatch(context.Background(), []string{"pager"}); err != nil {
		t.Fatalf("fresh dispatch inherited old cancellation: %v", err)
	}
	if got := gateway.Sent(); len(got) != 1 || got[0] != "pager" {
		t.Fatalf("unexpected sent targets: %#v", got)
	}
}
