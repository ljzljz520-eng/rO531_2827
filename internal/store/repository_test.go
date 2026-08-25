package store

import (
	"errors"
	"hospitalportal/internal/domain"
	"testing"
)

func TestMissingEntitiesUseSentinels(t *testing.T) {
	s, err := Open(t.TempDir() + "/db")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	if _, err = s.GetDepartment("missing"); !errors.Is(err, domain.ErrDepartmentNotFound) {
		t.Fatalf("got %v", err)
	}
	if _, err = s.GetShift("missing"); !errors.Is(err, domain.ErrShiftNotFound) {
		t.Fatalf("got %v", err)
	}
}
