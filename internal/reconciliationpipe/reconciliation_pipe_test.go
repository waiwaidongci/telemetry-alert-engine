package reconciliationpipe

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestConcurrentSourcesCloseAfterAllResults(t *testing.T) {
	coordinator := NewCoordinator(Producer{})
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	readingChannel, errorChannel := coordinator.Launch(ctx, map[string][]int{
		"alpha": {3, -1, 1},
		"beta":  {4, 2},
	})
	readings, failures := (Consumer{}).Collect(ctx, readingChannel, errorChannel)
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
