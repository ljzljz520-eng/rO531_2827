package domain

import (
	"testing"
	"time"
)

func TestShiftOverlap(t *testing.T) {
	base := time.Unix(1000, 0)
	a, err := NewDutyShift("a", "d", "u", "First", base, base.Add(4*time.Hour), base)
	if err != nil {
		t.Fatal(err)
	}
	b, err := NewDutyShift("b", "d", "u", "Second", base.Add(3*time.Hour), base.Add(6*time.Hour), base)
	if err != nil {
		t.Fatal(err)
	}
	if !a.Overlaps(b) {
		t.Fatal("expected overlap")
	}
}
