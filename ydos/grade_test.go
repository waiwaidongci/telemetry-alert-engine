package b_deliverygrade

import (
	"errors"
	"net/http"
	"testing"

	core "github.com/example/telemetry-alert/internal/notificationerrors"
)

func TestAbsentDeliveryRetainsSentinel(t *testing.T) {
	err := (core.Repository{}).Load("missing")
	if !errors.Is(err, core.ErrMissingAttempt) {
		t.Errorf("chain: %v", err)
	}
	kind := core.Classify(err)
	if kind != core.KindMissing {
		t.Errorf("kind=%q", kind)
	}
	if core.HTTPStatus(kind) != http.StatusNotFound {
		t.Errorf("status=%d", core.HTTPStatus(kind))
	}
	if core.ShouldRetry(kind, 0) {
		t.Error("missing delivery retried")
	}
}
