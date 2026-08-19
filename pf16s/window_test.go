package pf16s

import (
	"reflect"
	"testing"

	window "github.com/example/telemetry-alert/internal/metricwindow"
)

func clone(points []window.Point) []window.Point {
	result := make([]window.Point, len(points))
	for index, point := range points {
		result[index] = point
		result[index].Tags = make(map[string]string, len(point.Tags))
		for key, value := range point.Tags {
			result[index].Tags[key] = value
		}
	}
	return result
}

func TestFoxtrotWindow016(t *testing.T) {
	source := []window.Point{
		{Sequence: 1, Value: 3, Tags: map[string]string{"state": "cold"}},
		{Sequence: 2, Value: 11, Tags: map[string]string{"state": "hot"}},
		{Sequence: 3, Value: 15, Tags: map[string]string{"state": "hot"}},
	}
	original := clone(source)
	store := window.NewStore()
	cache := &window.Cache{}
	service := window.NewService(store, cache)
	published := service.Build("device-a", source, 10)
	published[0].Tags["state"] = "client-mutated"
	published[0].Value = -1
	service.Extend("device-a", window.Point{Sequence: 4, Value: 20, Tags: map[string]string{"state": "hot"}})
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
