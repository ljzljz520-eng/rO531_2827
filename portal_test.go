package hospitalportal

import (
	"errors"
	"hospitalportal/internal/account"
	"hospitalportal/internal/domain"
	"testing"
	"time"
)

func testNow() time.Time { return time.Date(2026, 1, 2, 9, 0, 0, 0, time.UTC) }
func seededPortal(t *testing.T) *Portal {
	t.Helper()
	p, err := Open(t.TempDir()+"/portal.db", testNow)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })
	d, err := p.Departments.Create("dep-er", "ER", "Emergency")
	if err != nil {
		t.Fatal(err)
	}
	_, err = p.Departments.UpdateDetails(d.ID, domain.DepartmentDetails{Location: "Building A", Phone: "1001", Email: "er@hospital.test", Services: []string{"triage", "emergency medicine"}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = p.Departments.Activate(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	admin, err := p.Accounts.Create(account.CreateCommand{ID: "admin-1", DisplayName: "Portal Admin", EmployeeNumber: "A001", Email: "admin@hospital.test", Role: domain.RoleAdministrator})
	if err != nil {
		t.Fatal(err)
	}
	_, err = p.Accounts.Activate(admin.ID)
	if err != nil {
		t.Fatal(err)
	}
	nurse, err := p.Accounts.Create(account.CreateCommand{ID: "nurse-1", DisplayName: "Duty Nurse", EmployeeNumber: "N001", Email: "nurse@hospital.test", Role: domain.RoleNurse, DepartmentID: d.ID})
	if err != nil {
		t.Fatal(err)
	}
	_, err = p.Accounts.Activate(nurse.ID)
	if err != nil {
		t.Fatal(err)
	}
	return p
}
func TestBusinessChain33(t *testing.T) {
	p := seededPortal(t)
	start := testNow().Add(time.Hour)
	_, err := p.ActivateAccountAndPublishShift("admin-1", "nurse-1", "shift-1", "Morning duty", start, start.Add(8*time.Hour))
	if !errors.Is(err, domain.ErrAccountState) {
		t.Fatalf("expected account state error, got %v", err)
	}
	shifts, listErr := p.Schedules.List("dep-er", "nurse-1", time.Time{}, time.Time{})
	if listErr != nil {
		t.Fatal(listErr)
	}
	if len(shifts) != 0 {
		t.Fatalf("workflow continued and created %d shifts", len(shifts))
	}
}
