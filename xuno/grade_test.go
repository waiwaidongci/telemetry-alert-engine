package a_ingestgrade

import (
	"context"
	"errors"
	"testing"
	"time"

	core "github.com/example/telemetry-alert/internal/ingestcontext"
)

func TestTelemetryAbortStopsEachStage(t *testing.T) {
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(core.NewRequest(cancelled).Context().Err(), context.Canceled) {
		t.Error("request lost cancellation")
	}
	started := time.Now()
	if err := (core.Gateway{}).Load(cancelled, make(chan struct{})); !errors.Is(err, context.Canceled) || time.Since(started) > 100*time.Millisecond {
		t.Errorf("gateway ignored cancellation: %v", err)
	}
	calls := 0
	if err := (core.Worker{}).Retry(cancelled, 3, func(context.Context) error { calls++; return core.ErrGatewayUnavailable }); !errors.Is(err, context.Canceled) || calls != 0 {
		t.Errorf("worker calls=%d err=%v", calls, err)
	}
	type key struct{}
	ctx := context.WithValue(context.Background(), key{}, "kept")
	ready := make(chan struct{})
	close(ready)
	seen := false
	if err := core.NewService().Ingest(ctx, ready, func(got context.Context) error { seen = got.Value(key{}) == "kept"; return nil }); err != nil || !seen {
		t.Errorf("service context seen=%v err=%v", seen, err)
	}
}
