package notificationerrors

import (
	"errors"
	"net/http"
	"testing"
)

func TestAbsentDeliveryPreservesSentinelAndStopsQueue(t *testing.T) {
	err := (Repository{}).Load("missing")
	if !errors.Is(err, ErrMissingAttempt) {
		t.Errorf("repository broke sentinel chain: %v", err)
	}
	if got := Classify(err); got != KindMissing {
		t.Errorf("classifier returned %q", got)
	}
	if got := HTTPStatus(KindMissing); got != http.StatusNotFound {
		t.Errorf("handler returned %d", got)
	}
	if ShouldRetry(KindMissing, 0) {
		t.Error("missing attempt was scheduled for retry")
	}
}
