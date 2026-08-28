package schedule

import (
	"hospitalportal/internal/domain"
	"sort"
	"time"
)

type CoverageWindow struct {
	StartAt, EndAt          time.Time
	DoctorCount, NurseCount int
	Uncovered               bool
}

func CalculateCoverage(shifts []domain.DutyShift, accounts []domain.UserAccount, from, to time.Time, step time.Duration) []CoverageWindow {
	roles := map[string]domain.Role{}
	for _, a := range accounts {
		roles[a.ID] = a.Role
	}
	out := []CoverageWindow{}
	if step <= 0 || !to.After(from) {
		return out
	}
	for start := from; start.Before(to); start = start.Add(step) {
		end := start.Add(step)
		if end.After(to) {
			end = to
		}
		w := CoverageWindow{StartAt: start, EndAt: end}
		for _, shift := range shifts {
			if shift.Status != domain.ShiftPublished || !shift.StartAt.Before(end) || !start.Before(shift.EndAt) {
				continue
			}
			switch roles[shift.AccountID] {
			case domain.RoleDoctor:
				w.DoctorCount++
			case domain.RoleNurse:
				w.NurseCount++
			}
		}
		w.Uncovered = w.DoctorCount == 0 || w.NurseCount == 0
		out = append(out, w)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartAt.Before(out[j].StartAt) })
	return out
}
