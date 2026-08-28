package hospitalportal

import (
	"hospitalportal/internal/domain"
	"testing"
)

func TestDepartmentWorkflow(t *testing.T) {
	p, err := Open(t.TempDir()+"/portal.db", testNow)
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	d, err := p.Departments.Create("dep-card", "CARD", "Cardiology")
	if err != nil {
		t.Fatal(err)
	}
	d, err = p.Departments.UpdateDetails(d.ID, domain.DepartmentDetails{Location: "Building C", Phone: "2002", Email: "card@hospital.test", Services: []string{"ECG", "consultation", "ECG"}})
	if err != nil {
		t.Fatal(err)
	}
	d, err = p.Departments.Activate(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := p.Departments.Get(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Status != domain.DepartmentActive || len(loaded.Services) != 2 {
		t.Fatalf("unexpected department: %#v", loaded)
	}
}
