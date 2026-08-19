package reconciliationpipe

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestConcurrentSourcesCloseAfterAllResults(t *testing.T) {
	t.Run("error bus honors capacity and wrapping", func(t *testing.T) {
		bus := NewErrorBus(2)
		reported := make(chan struct{})
		go func() {
			bus.Report("alpha", ErrInvalidReading)
			bus.Report("beta", ErrInvalidReading)
			close(reported)
		}()
		select {
		case <-reported:
		case <-time.After(20 * time.Millisecond):
			<-bus.Channel()
			<-reported
			t.Error("error bus blocked before reaching configured capacity")
		}
		bus.Close()
		count := 0
		for err := range bus.Channel() {
			count++
			if !errors.Is(err, ErrInvalidReading) {
				t.Errorf("error identity was lost: %v", err)
			}
		}
		if count != 2 {
			t.Errorf("unexpected error count %d", count)
		}
	})

	t.Run("producer continues after invalid reading", func(t *testing.T) {
		start := make(chan struct{})
		close(start)
		output := make(chan Reading, 2)
		bus := NewErrorBus(2)
		(Producer{}).Run(context.Background(), "alpha", []int{-1, 7}, output, bus, start)
		close(output)
		bus.Close()
		if got := len(output); got != 1 {
			t.Errorf("producer lost valid readings after an error: %d", got)
		}
	})

	t.Run("consumer drains both channels", func(t *testing.T) {
		output := make(chan Reading, 1)
		failures := make(chan error, 1)
		output <- Reading{Source: "alpha", Value: 1}
		failures <- ErrInvalidReading
		close(output)
		close(failures)
		readings, errs := (Consumer{}).Collect(context.Background(), output, failures)
		if len(readings) != 1 || len(errs) != 1 {
			t.Errorf("consumer did not drain both channels: readings=%d errors=%d", len(readings), len(errs))
		}
	})

	start := make(chan struct{})
	coordinator := NewCoordinator(Producer{}, start)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	readingChannel, errorChannel := coordinator.Launch(ctx, map[string][]int{
		"alpha": {3, -1, 1},
		"beta":  {4, 2},
	})
	type result struct {
		readings []Reading
		failures []error
	}
	resultChannel := make(chan result, 1)
	go func() {
		readings, failures := (Consumer{}).Collect(ctx, readingChannel, errorChannel)
		resultChannel <- result{readings: readings, failures: failures}
	}()
	close(start)
	resultValue := <-resultChannel
	readings, failures := resultValue.readings, resultValue.failures
	want := []Reading{{Source: "alpha", Value: 1}, {Source: "alpha", Value: 3}, {Source: "beta", Value: 2}, {Source: "beta", Value: 4}}
	if !reflect.DeepEqual(readings, want) {
		t.Fatalf("incomplete readings: %#v", readings)
	}
	if len(failures) != 1 || !errors.Is(failures[0], ErrInvalidReading) {
		t.Fatalf("unexpected failures: %#v", failures)
	}
	if errors.Is(failures[0], context.DeadlineExceeded) {
		t.Fatal("pipeline did not close before deadline")
	}
}
