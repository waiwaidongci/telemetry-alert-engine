package windowcache

import "testing"

func TestSnapshotsDoNotShareMutableMaps(t *testing.T) {
	s := NewStore()
	s.Put("cpu", 10)
	snap := s.Snapshot()
	s.Put("cpu", 20)
	if snap["cpu"] != 10 {
		t.Fatalf("snapshot changed to %d", snap["cpu"])
	}
	c := &Cache{}
	source := map[string]int{"mem": 7}
	c.Store(source)
	source["mem"] = 9
	if c.Read("mem") != 7 {
		t.Fatalf("cache changed to %d", c.Read("mem"))
	}
	raw := map[string]int{"disk": -1}
	Normalize(raw)
	if raw["disk"] != -1 {
		t.Fatalf("normalize mutated input: %#v", raw)
	}
	left := map[string]int{"a": 1}
	Merge(left, map[string]int{"b": 2})
	if len(left) != 1 {
		t.Fatalf("merge mutated left: %#v", left)
	}
}
