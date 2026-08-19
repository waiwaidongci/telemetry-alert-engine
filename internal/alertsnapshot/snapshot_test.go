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
	store.Upsert(Rule{ID: "rule-disabled", Active: false, Labels: map[string]string{"site": "west"}})
	first := service.Refresh()

	ready := make(chan struct{})
	done := make(chan struct{}, 2)
	var wait sync.WaitGroup
	wait.Add(2)
	writerRule := Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "south"}}
	go func() {
		defer wait.Done()
		<-ready
		for index := 0; index < 200; index++ {
			store.Upsert(Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "south"}})
		}
		done <- struct{}{}
	}()
	go func() {
		defer wait.Done()
		<-ready
		for index := 0; index < 200; index++ {
			current := service.Refresh()
			if len(current) != 1 && index == 0 {
				t.Errorf("unexpected snapshot size %d", len(current))
			}
		}
		done <- struct{}{}
	}()
	close(ready)
	<-done
	<-done
	wait.Wait()
	writerDone := make(chan struct{}, 1)
	writer.Apply(writerRule, ready, writerDone)
	<-writerDone

	if len(first) != 1 || first[0].ID != "rule-a" {
		t.Fatalf("refresh included inactive or unordered rules: %#v", first)
	}
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
	if _, ok := writerRule.Labels["writer"]; ok {
		t.Fatal("writer mutated caller labels")
	}
}
