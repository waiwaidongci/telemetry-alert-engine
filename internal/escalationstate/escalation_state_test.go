package escalationstate

import (
	"reflect"
	"testing"
)

func TestRetryRecoveryReachesDeliveredState(t *testing.T) {
	incident := Incident{ID: "incident-a", State: StateFailed, History: []State{StateQueued, StateFailed}}
	worker := NewWorker(Machine{})
	recovered, err := worker.Recover(incident, func() error { return nil })
	if err != nil {
		t.Fatalf("recover escalation: %v", err)
	}
	if recovered.State != StateDelivered {
		t.Fatalf("unexpected final state %q", recovered.State)
	}
	wantHistory := []State{StateQueued, StateFailed, StateRetrying, StateDelivered}
	if !reflect.DeepEqual(recovered.History, wantHistory) {
		t.Fatalf("unexpected history: %#v", recovered.History)
	}
	query := Query{}
	if len(query.InProgress([]Incident{recovered})) != 0 {
		t.Fatal("delivered incident remained in progress")
	}
	if len(query.Delivered([]Incident{recovered})) != 1 {
		t.Fatal("delivered incident missing from terminal query")
	}
}

func TestRetryingIncidentRemainsVisible(t *testing.T) {
	incident := Incident{ID: "incident-b", State: StateRetrying}
	if got := (Query{}).InProgress([]Incident{incident}); len(got) != 1 {
		t.Fatalf("retrying incident disappeared: %#v", got)
	}
}
