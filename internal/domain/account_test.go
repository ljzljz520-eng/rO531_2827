package domain

import (
	"testing"
	"time"
)

func TestAccountLifecycle(t *testing.T) {
	a, err := NewUserAccount("u1", "Name", "E1", "n@h.test", RoleNurse, "d1", time.Unix(1, 0))
	if err != nil {
		t.Fatal(err)
	}
	a, err = a.Activate(time.Unix(2, 0))
	if err != nil {
		t.Fatal(err)
	}
	if !a.Can("schedule.manage") {
		t.Fatal("active nurse should manage schedule")
	}
	a, err = a.Suspend(time.Unix(3, 0))
	if err != nil {
		t.Fatal(err)
	}
	if a.Can("schedule.read") {
		t.Fatal("suspended account retained permission")
	}
}
