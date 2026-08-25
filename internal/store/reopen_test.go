package store

import (
	"hospitalportal/internal/domain"
	"testing"
	"time"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := t.TempDir() + "/portal.db"
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	a, err := domain.NewUserAccount("doctor-1", "Doctor One", "D001", "doctor@hospital.test", domain.RoleDoctor, "dep-1", time.Unix(100, 0))
	if err != nil {
		t.Fatal(err)
	}
	if err = s.CreateAccount(a); err != nil {
		t.Fatal(err)
	}
	if err = s.Close(); err != nil {
		t.Fatal(err)
	}
	s, err = Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	loaded, err := s.GetAccount(a.ID)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.EmployeeNumber != a.EmployeeNumber {
		t.Fatalf("got %#v", loaded)
	}
}
