package c_windowgrade

import (
	core "github.com/example/telemetry-alert/internal/querysnapshot"
	"testing"
)

func TestPublishedWindowsOwnTheirArrays(t *testing.T) {
	original := []core.Point{{ID: "old"}, {ID: "live", Active: true}}
	store := &core.Store{}
	store.Save(original)
	original[0].ID = "mutated"
	if store.Snapshot()[0].ID != "old" {
		t.Error("store aliases caller")
	}
	input := []core.Point{{ID: "drop"}, {ID: "keep", Active: true}}
	filtered := core.FilterActive(input)
	filtered[0].ID = "changed"
	if input[0].ID != "drop" || input[1].ID != "keep" {
		t.Error("filter rewrote input")
	}
	cache := &core.Cache{}
	values := []core.Point{{ID: "snapshot"}}
	cache.Publish(values)
	values[0].ID = "later"
	if cache.Read()[0].ID != "snapshot" {
		t.Error("cache aliases publisher")
	}
	base := make([]core.Point, 1, 4)
	base[0].ID = "base"
	built := (core.Service{}).Build(base, core.Point{ID: "extra"})
	built[0].ID = "changed"
	if base[0].ID != "base" {
		t.Error("service aliases base")
	}
}

func TestVacantPublishedWindowReturnsEmpty(t *testing.T) {
	if len((&core.Cache{}).Read()) != 0 {
		t.Error("empty cache returned data")
	}
}
