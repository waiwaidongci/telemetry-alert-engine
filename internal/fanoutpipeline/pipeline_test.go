package fanoutpipeline

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestNotificationFanoutWaitsForAllTargets(t *testing.T) {
	results := Produce(context.Background(), []string{"a", "b"}, func(string) error { return nil })
	collected, err := Consume(context.Background(), results)
	if err != nil || len(collected) != 2 {
		t.Errorf("producer/consumer lifecycle failed: len=%d err=%v", len(collected), err)
	}

	start := make(chan struct{})
	done := make(chan struct{})
	var ran atomic.Int32
	Launch(start, []func(){func() { ran.Add(1) }, func() { ran.Add(1) }}, done)
	select {
	case <-done:
		t.Error("coordinator completed before workers started")
	default:
	}
	close(start)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("coordinator did not finish")
	}
	if ran.Load() != 2 {
		t.Errorf("only %d workers ran", ran.Load())
	}

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Consume(cancelled, make(chan Result)); !errors.Is(err, context.Canceled) {
		t.Errorf("consumer ignored cancellation: %v", err)
	}
	if cap(NewErrorChannel(3)) < 3 {
		t.Error("error channel cannot accept one error per worker")
	}
}
