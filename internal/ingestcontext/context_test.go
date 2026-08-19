package ingestcontext

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRequestCancellationStopsTelemetryRetries(t *testing.T) {
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(NewRequest(cancelled).Context().Err(), context.Canceled) {
		t.Error("request lost its cancellation")
	}
	if NewRequest(context.Background()).Context().Err() != nil {
		t.Error("a cancelled request polluted the next request")
	}

	started := time.Now()
	if err := (Gateway{}).Load(cancelled, make(chan struct{})); !errors.Is(err, context.Canceled) || time.Since(started) > 100*time.Millisecond {
		t.Errorf("gateway ignored cancellation: %v", err)
	}

	calls := 0
	err := (Worker{}).Retry(cancelled, 3, func(context.Context) error { calls++; return ErrGatewayUnavailable })
	if !errors.Is(err, context.Canceled) || calls != 0 {
		t.Errorf("worker retried after cancellation: calls=%d err=%v", calls, err)
	}

	type requestKey struct{}
	requestContext := context.WithValue(context.Background(), requestKey{}, "request-42")
	ready := make(chan struct{})
	close(ready)
	seenRequest := false
	err = NewService().Ingest(requestContext, ready, func(ctx context.Context) error {
		seenRequest = ctx.Value(requestKey{}) == "request-42"
		return nil
	})
	if err != nil || !seenRequest {
		t.Errorf("service replaced request context: seen=%v err=%v", seenRequest, err)
	}
}
