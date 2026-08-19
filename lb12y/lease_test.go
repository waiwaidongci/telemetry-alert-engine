package lb12y

import (
	"errors"
	"reflect"
	"testing"

	lease "github.com/example/telemetry-alert/internal/exportlease"
)

func TestBravoLease012(t *testing.T) {
	pool := lease.NewPool(2)
	repository := &lease.Repository{}
	service := lease.NewService(lease.NewBatch(pool, repository))
	err := service.Execute([]lease.Item{{ID: "a"}, {ID: "b"}, {ID: "bad", Reject: true}, {ID: "c"}})
	if !errors.Is(err, lease.ErrRejectedPayload) {
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
