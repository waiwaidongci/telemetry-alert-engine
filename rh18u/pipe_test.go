package rh18u

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	pipe "github.com/example/telemetry-alert/internal/reconciliationpipe"
)

func TestHotelPipeline018(t *testing.T) {
	t.Run("error bus", func(t *testing.T) {
		bus := pipe.NewErrorBus(2)
		reported := make(chan struct{})
		go func() {
			bus.Report("alpha", pipe.ErrInvalidReading)
			bus.Report("beta", pipe.ErrInvalidReading)
			close(reported)
		}()
		select {
		case <-reported:
		case <-time.After(20 * time.Millisecond):
			<-bus.Channel()
			<-reported
			t.Error("error bus blocked before configured capacity")
		}
		bus.Close()
		count := 0
		for err := range bus.Channel() {
			count++
			if !errors.Is(err, pipe.ErrInvalidReading) {
				t.Errorf("error identity was lost: %v", err)
			}
		}
		if count != 2 {
			t.Errorf("unexpected error count %d", count)
		}
	})

	t.Run("producer", func(t *testing.T) {
		start := make(chan struct{})
		close(start)
		output := make(chan pipe.Reading, 2)
		bus := pipe.NewErrorBus(2)
		(pipe.Producer{}).Run(context.Background(), "alpha", []int{-1, 7}, output, bus, start)
		close(output)
		bus.Close()
		if len(output) != 1 {
			t.Errorf("producer lost valid reading")
		}
	})

	t.Run("consumer", func(t *testing.T) {
		output := make(chan pipe.Reading, 1)
		failures := make(chan error, 1)
		output <- pipe.Reading{Source: "alpha", Value: 1}
		failures <- pipe.ErrInvalidReading
		close(output)
		close(failures)
		readings, errs := (pipe.Consumer{}).Collect(context.Background(), output, failures)
		if len(readings) != 1 || len(errs) != 1 {
			t.Errorf("consumer did not drain both channels")
		}
	})

	start := make(chan struct{})
	coordinator := pipe.NewCoordinator(pipe.Producer{}, start)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	readingsChannel, errorsChannel := coordinator.Launch(ctx, map[string][]int{
		"alpha": {3, -1, 1},
		"beta":  {4, 2},
	})
	type result struct {
		readings []pipe.Reading
		failures []error
	}
	results := make(chan result, 1)
	go func() {
		readings, failures := (pipe.Consumer{}).Collect(ctx, readingsChannel, errorsChannel)
		results <- result{readings: readings, failures: failures}
	}()
	close(start)
	got := <-results
	want := []pipe.Reading{{Source: "alpha", Value: 1}, {Source: "alpha", Value: 3}, {Source: "beta", Value: 2}, {Source: "beta", Value: 4}}
	if !reflect.DeepEqual(got.readings, want) {
		t.Fatalf("incomplete readings: %#v", got.readings)
	}
	if len(got.failures) != 1 || !errors.Is(got.failures[0], pipe.ErrInvalidReading) {
		t.Fatalf("unexpected failures: %#v", got.failures)
	}
}
