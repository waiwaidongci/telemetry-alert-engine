package devicequery

import (
	"errors"
	"testing"
)

func TestMissingDeviceKeepsNotFoundContract(t *testing.T) {
	err := RepositoryError(ErrNotFound)
	if !errors.Is(err, ErrNotFound) {
		t.Fatal("not-found identity was lost")
	}
	kind := Classify(err)
	if kind != NotFound {
		t.Fatalf("kind=%q", kind)
	}
	status := HTTPStatus(kind)
	if status != 404 {
		t.Fatalf("status=%d", status)
	}
	if ShouldRetry(status) {
		t.Fatal("404 must not be retried")
	}
}
