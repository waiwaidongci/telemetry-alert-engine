package e_fanoutgrade

import (
	"context"
	"errors"
	core "github.com/example/telemetry-alert/internal/fanoutpipeline"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAllNotificationTargetsFinishBeforeDone(t *testing.T) {
	barrier := make(chan struct{})
	var parallel sync.WaitGroup
	parallel.Add(2)
	exercise := func() {
		<-barrier
		stream := core.Produce(context.Background(), []string{"left", "right"}, func(string) error { return nil })
		_, _ = core.Consume(context.Background(), stream)
	}
	go func() { defer parallel.Done(); exercise() }()
	go func() { defer parallel.Done(); exercise() }()
	close(barrier)
	parallel.Wait()

	results := core.Produce(context.Background(), []string{"a", "b"}, func(string) error { return nil })
	got, err := core.Consume(context.Background(), results)
	if err != nil || len(got) != 2 {
		t.Errorf("len=%d err=%v", len(got), err)
	}
	start := make(chan struct{})
	done := make(chan struct{})
	var ran atomic.Int32
	core.Launch(start, []func(){func() { ran.Add(1) }, func() { ran.Add(1) }}, done)
	select {
	case <-done:
		t.Error("done closed early")
	default:
	}
	close(start)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("workers stuck")
	}
	if ran.Load() != 2 {
		t.Errorf("ran=%d", ran.Load())
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := core.Consume(ctx, make(chan core.Result)); !errors.Is(err, context.Canceled) {
		t.Errorf("cancel=%v", err)
	}
	if cap(core.NewErrorChannel(3)) < 3 {
		t.Error("error channel too small")
	}
}
