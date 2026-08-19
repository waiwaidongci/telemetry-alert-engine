package metricwindow

import (
	"reflect"
	"testing"
)

func TestWindowCopiesAcrossFilterStoreAndCache(t *testing.T) {
	source := []Point{
		{Sequence: 1, Value: 3, Tags: map[string]string{"state": "cold"}},
		{Sequence: 2, Value: 11, Tags: map[string]string{"state": "hot"}},
		{Sequence: 3, Value: 15, Tags: map[string]string{"state": "hot"}},
	}
	original := clonePoints(source)
	store := NewStore()
	cache := &Cache{}
	service := NewService(store, cache)
	published := service.Build("device-a", source, 10)
	published[0].Tags["state"] = "client-mutated"
	published[0].Value = -1
	service.Extend("device-a", Point{Sequence: 4, Value: 20, Tags: map[string]string{"state": "hot"}})

	if !reflect.DeepEqual(source, original) {
		t.Fatalf("filter changed source slice: %#v", source)
	}
	stored := store.Load("device-a")
	if len(stored) != 3 || stored[0].Value != 11 || stored[0].Tags["state"] != "hot" {
		t.Fatalf("stored window was aliased: %#v", stored)
	}
	current := cache.Current()
	current[0].Tags["state"] = "another-client"
	if cache.Current()[0].Tags["state"] != "hot" {
		t.Fatal("cache returned internal backing data")
	}
}
