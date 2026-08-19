package exportlease

import (
	"errors"
	"reflect"
	"testing"
)

func TestBatchReleasesResourcesAndPreservesFailure(t *testing.T) {
	pool := NewPool(2)
	repository := &Repository{}
	service := NewService(NewBatch(pool, repository))
	err := service.Execute([]Item{{ID: "a"}, {ID: "b"}, {ID: "bad", Reject: true}, {ID: "c"}})
	if !errors.Is(err, ErrRejectedPayload) {
		t.Fatalf("expected rejected payload, got %v", err)
	}
	if pool.Open() != 0 {
		t.Fatalf("resources still open: %d", pool.Open())
	}
	if service.Reserved() {
		t.Fatal("service reservation was not released")
	}
	if got := repository.Committed(); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("unexpected committed exports: %#v", got)
	}
}
