package tj20w

import (
	"context"
	"errors"
	"testing"
	"time"

	dispatch "github.com/example/telemetry-alert/internal/dispatchcontext"
)

func TestJulietDispatch020(t *testing.T) {
	t.Run("timeout cause", TestJulietTimeout020)
	t.Run("request isolation", TestJulietIsolation020)
	t.Run("parent cancellation", TestJulietParent020)
}

func TestJulietTimeout020(t *testing.T) {
	gateway := dispatch.NewGateway(40 * time.Millisecond)
	service := dispatch.NewService(dispatch.NewWorker(gateway), 5*time.Millisecond)
	started := time.Now()
	err := service.Dispatch(context.Background(), []string{"pager", "webhook", "email"})
	if !errors.Is(err, dispatch.ErrDispatchDeadline) {
		t.Fatalf("timeout cause was lost: %v", err)
	}
	if time.Since(started) > 30*time.Millisecond {
		t.Fatal("fanout ignored dispatch timeout")
	}
	time.Sleep(45 * time.Millisecond)
	if len(gateway.Sent()) != 0 {
		t.Fatal("cancelled deliveries continued")
	}
}

func TestJulietIsolation020(t *testing.T) {
	gateway := dispatch.NewGateway(30 * time.Millisecond)
	service := dispatch.NewService(dispatch.NewWorker(gateway), 5*time.Millisecond)
	if err := service.Dispatch(context.Background(), []string{"slow"}); !errors.Is(err, dispatch.ErrDispatchDeadline) {
		t.Fatalf("first dispatch did not time out: %v", err)
	}
	gateway.SetDelay(time.Millisecond)
	if err := service.Dispatch(context.Background(), []string{"pager"}); err != nil {
		t.Fatalf("fresh dispatch inherited cancellation: %v", err)
	}
	sent := gateway.Sent()
	if len(sent) != 1 || sent[0] != "pager" {
		t.Fatalf("unexpected sent targets: %#v", sent)
	}
}

func TestJulietParent020(t *testing.T) {
	gateway := dispatch.NewGateway(40 * time.Millisecond)
	service := dispatch.NewService(dispatch.NewWorker(gateway), 100*time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	err := service.Dispatch(ctx, []string{"pager", "email"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("parent cancellation was lost: %v", err)
	}
	if time.Since(started) > 20*time.Millisecond {
		t.Fatal("cancelled parent did not stop promptly")
	}
	if len(gateway.Sent()) != 0 {
		t.Fatal("cancelled parent still sent targets")
	}
}
