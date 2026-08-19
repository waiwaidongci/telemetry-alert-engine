package querysnapshot

import "testing"

func TestWindowCopiesRemainIndependent(t *testing.T) {
	original := []Point{{ID: "old", Active: false}, {ID: "live", Active: true}}
	store := &Store{}
	store.Save(original)
	original[0].ID = "mutated"
	if got := store.Snapshot()[0].ID; got != "old" {
		t.Errorf("store retained caller array: %q", got)
	}

	filteredInput := []Point{{ID: "drop"}, {ID: "keep", Active: true}}
	filtered := FilterActive(filteredInput)
	filtered[0].ID = "changed"
	if filteredInput[0].ID != "drop" || filteredInput[1].ID != "keep" {
		t.Errorf("filter rewrote input: %#v", filteredInput)
	}

	cache := &Cache{}
	published := []Point{{ID: "snapshot", Active: true}}
	cache.Publish(published)
	published[0].ID = "later"
	if got := cache.Read()[0].ID; got != "snapshot" {
		t.Errorf("cache snapshot changed: %q", got)
	}

	base := make([]Point, 1, 4)
	base[0] = Point{ID: "base"}
	built := (Service{}).Build(base, Point{ID: "extra"})
	built[0].ID = "rewritten"
	if base[0].ID != "base" {
		t.Errorf("service result aliased base: %q", base[0].ID)
	}
}

func TestVacantWindowReadIsSafe(t *testing.T) {
	if got := (&Cache{}).Read(); len(got) != 0 {
		t.Errorf("empty cache returned %d points", len(got))
	}
}
