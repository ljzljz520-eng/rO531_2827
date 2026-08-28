package hospitalportal

import (
	"hospitalportal/internal/domain"
	"hospitalportal/internal/schedule"
	"testing"
	"time"
)

func TestDutyWorkflow(t *testing.T) {
	p := seededPortal(t)
	start := testNow().Add(24 * time.Hour)
	shift, err := p.Schedules.Create(schedule.CreateCommand{ID: "duty-1", DepartmentID: "dep-er", AccountID: "nurse-1", Title: "Day duty", StartAt: start, EndAt: start.Add(8 * time.Hour)})
	if err != nil {
		t.Fatal(err)
	}
	shift, err = p.Schedules.Publish(shift.ID)
	if err != nil {
		t.Fatal(err)
	}
	items, err := p.Schedules.List("dep-er", "nurse-1", start, start.Add(12*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].Status != domain.ShiftPublished {
		t.Fatalf("unexpected shifts: %#v", items)
	}
}
