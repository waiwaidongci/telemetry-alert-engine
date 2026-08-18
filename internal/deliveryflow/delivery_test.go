package deliveryflow

import "testing"

func TestRetrySuccessReachesSentAndActiveQueryIncludesRetrying(t *testing.T) {
	s := NewService()
	d := &Delivery{ID: "d1", Status: Queued}
	if !s.Transition(d, Retrying) {
		t.Fatal("queued delivery was not retried")
	}
	if d.Status.Terminal() {
		t.Fatal("retrying delivery was treated as terminal")
	}
	if !RetrySucceeded(s, d) {
		t.Fatal("successful retry was rejected")
	}
	if d.Status != Sent {
		t.Fatalf("status = %q, want sent", d.Status)
	}
	items := Active([]Delivery{{ID: "d2", Status: Retrying}, {ID: "d3", Status: Sent}})
	if len(items) != 1 || items[0].ID != "d2" {
		t.Fatalf("active items = %#v", items)
	}
}
