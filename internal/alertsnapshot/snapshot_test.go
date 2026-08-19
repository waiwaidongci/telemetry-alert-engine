package alertsnapshot

import (
	"sync"
	"testing"
)

func TestSnapshotIsolationConcurrentUpdates(t *testing.T) {
	store := NewStore()
	cache := &Cache{}
	service := NewService(store, cache)
	writer := NewWriter(store)
	original := Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "north"}}
	store.Upsert(original)
	first := service.Refresh()

	ready := make(chan struct{})
	done := make(chan struct{}, 2)
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		writer.Apply(Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "south"}}, ready, done)
	}()
	go func() {
		defer wait.Done()
		<-ready
		for index := 0; index < 200; index++ {
			current := service.Refresh()
			if len(current) != 1 {
				t.Errorf("unexpected snapshot size %d", len(current))
				return
			}
		}
		done <- struct{}{}
	}()
	close(ready)
	<-done
	<-done
	wait.Wait()

	if got := first[0].Labels["site"]; got != "north" {
		t.Fatalf("published snapshot changed to %q", got)
	}
	first[0].Labels["site"] = "client-change"
	if got := cache.Current()[0].Labels["site"]; got == "client-change" {
		t.Fatal("cache exposed mutable labels")
	}
	if got := original.Labels["site"]; got != "north" {
		t.Fatalf("caller input was mutated to %q", got)
	}
}
