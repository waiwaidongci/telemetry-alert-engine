package ka11x

import (
	"sync"
	"testing"

	snapshot "github.com/example/telemetry-alert/internal/alertsnapshot"
)

func TestKappaSnapshot011(t *testing.T) {
	store := snapshot.NewStore()
	cache := &snapshot.Cache{}
	service := snapshot.NewService(store, cache)
	writer := snapshot.NewWriter(store)
	original := snapshot.Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "north"}}
	store.Upsert(original)
	store.Upsert(snapshot.Rule{ID: "rule-disabled", Active: false, Labels: map[string]string{"site": "west"}})
	first := service.Refresh()

	start := make(chan struct{})
	done := make(chan struct{}, 2)
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		<-start
		for index := 0; index < 200; index++ {
			store.Upsert(snapshot.Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "south"}})
		}
		done <- struct{}{}
	}()
	go func() {
		defer wait.Done()
		<-start
		for index := 0; index < 200; index++ {
			current := service.Refresh()
			if len(current) != 1 && index == 0 {
				t.Errorf("unexpected active snapshot size %d", len(current))
			}
		}
		done <- struct{}{}
	}()
	close(start)
	<-done
	<-done
	wait.Wait()

	writerRule := snapshot.Rule{ID: "rule-a", Active: true, Labels: map[string]string{"site": "south"}}
	writerDone := make(chan struct{}, 1)
	writer.Apply(writerRule, start, writerDone)
	<-writerDone
	if len(first) != 1 || first[0].ID != "rule-a" {
		t.Fatalf("inactive rules escaped into snapshot: %#v", first)
	}
	first[0].Labels["site"] = "client-change"
	if cache.Current()[0].Labels["site"] == "client-change" {
		t.Fatal("cache exposed mutable labels")
	}
	if original.Labels["site"] != "north" {
		t.Fatal("store mutated caller input")
	}
	if _, exists := writerRule.Labels["writer"]; exists {
		t.Fatal("writer mutated caller labels")
	}
}
