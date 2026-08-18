package exportbatch

import (
	"errors"
	"testing"
)

func TestExportResourcesAndErrors(t *testing.T) {
	t.Run("batch closes each handle", func(t *testing.T) {
		tr := &Tracker{}
		ExportMany(tr, 8)
		if tr.Peak > 1 || tr.Open != 0 {
			t.Fatalf("peak=%d open=%d", tr.Peak, tr.Open)
		}
	})
	t.Run("business error is preserved", func(t *testing.T) {
		if !errors.Is(Persist(&Tx{}, ErrBusiness), ErrBusiness) {
			t.Fatal("business error was lost")
		}
	})
	t.Run("validation releases handle", func(t *testing.T) {
		tr := &Tracker{}
		if !errors.Is(ValidateExport(tr, true), ErrBusiness) {
			t.Fatal("missing validation error")
		}
		if tr.Open != 0 {
			t.Fatalf("open=%d", tr.Open)
		}
	})
	t.Run("close is idempotent", func(t *testing.T) {
		tr := &Tracker{}
		r := tr.Acquire()
		r.Close()
		r.Close()
		if tr.Open != 0 {
			t.Fatalf("open=%d", tr.Open)
		}
	})
}
