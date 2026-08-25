package schedule

import (
	"fmt"
	"hospitalportal/internal/domain"
	"hospitalportal/internal/store"
	"time"
)

type Service struct {
	store *store.Store
	now   func() time.Time
}

func New(s *store.Store, now func() time.Time) *Service { return &Service{store: s, now: now} }

type CreateCommand struct {
	ID, DepartmentID, AccountID, Title string
	StartAt, EndAt                     time.Time
}

func (s *Service) Create(c CreateCommand) (domain.DutyShift, error) {
	d, err := s.store.GetDepartment(c.DepartmentID)
	if err != nil {
		return domain.DutyShift{}, err
	}
	if d.Status != domain.DepartmentActive {
		return domain.DutyShift{}, fmt.Errorf("%w: department inactive", domain.ErrShiftConflict)
	}
	a, err := s.store.GetAccount(c.AccountID)
	if err != nil {
		return domain.DutyShift{}, err
	}
	if a.Status != domain.AccountActive || a.DepartmentID != c.DepartmentID {
		return domain.DutyShift{}, fmt.Errorf("%w: account unavailable", domain.ErrShiftConflict)
	}
	v, err := domain.NewDutyShift(c.ID, c.DepartmentID, c.AccountID, c.Title, c.StartAt, c.EndAt, s.now())
	if err != nil {
		return v, err
	}
	existing, err := s.store.ListShifts(c.DepartmentID, c.AccountID, c.StartAt, c.EndAt)
	if err != nil {
		return v, err
	}
	for _, other := range existing {
		if v.Overlaps(other) {
			return v, domain.ErrShiftConflict
		}
	}
	return v, s.store.CreateShift(v)
}
func (s *Service) Publish(id string) (domain.DutyShift, error) {
	v, err := s.store.GetShift(id)
	if err != nil {
		return v, err
	}
	v, err = v.Publish(s.now())
	if err != nil {
		return v, err
	}
	return v, s.store.PutShift(v)
}
func (s *Service) Cancel(id, reason string) (domain.DutyShift, error) {
	v, err := s.store.GetShift(id)
	if err != nil {
		return v, err
	}
	v, err = v.Cancel(reason, s.now())
	if err != nil {
		return v, err
	}
	return v, s.store.PutShift(v)
}
func (s *Service) List(department, account string, from, to time.Time) ([]domain.DutyShift, error) {
	return s.store.ListShifts(department, account, from, to)
}
