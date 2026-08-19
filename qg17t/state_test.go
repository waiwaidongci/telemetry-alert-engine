package qg17t

import (
	"reflect"
	"testing"

	state "github.com/example/telemetry-alert/internal/escalationstate"
)

func TestGolfState017(t *testing.T) {
	t.Run("recovery", TestGolfRecovery017)
	t.Run("visibility", TestGolfVisibility017)
}

func TestGolfRecovery017(t *testing.T) {
	incident := state.Incident{ID: "incident-a", State: state.StateFailed, History: []state.State{state.StateQueued, state.StateFailed}}
	recovered, err := state.NewWorker(state.Machine{}).Recover(incident, func() error { return nil })
	if err != nil {
		t.Fatalf("recover escalation: %v", err)
	}
	if recovered.State != state.StateDelivered {
		t.Fatalf("unexpected final state %q", recovered.State)
	}
	want := []state.State{state.StateQueued, state.StateFailed, state.StateRetrying, state.StateDelivered}
	if !reflect.DeepEqual(recovered.History, want) {
		t.Fatalf("unexpected history: %#v", recovered.History)
	}
	query := state.Query{}
	if len(query.InProgress([]state.Incident{recovered})) != 0 || len(query.Delivered([]state.Incident{recovered})) != 1 {
		t.Fatal("delivered incident is in the wrong query set")
	}
}

func TestGolfVisibility017(t *testing.T) {
	incident := state.Incident{ID: "incident-b", State: state.StateRetrying}
	if got := (state.Query{}).InProgress([]state.Incident{incident}); len(got) != 1 {
		t.Fatalf("retrying incident disappeared: %#v", got)
	}
}
