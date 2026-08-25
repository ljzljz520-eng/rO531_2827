package department

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
func (s *Service) Create(id, code, name string) (domain.Department, error) {
	d, err := domain.NewDepartment(id, code, name, s.now())
	if err != nil {
		return d, err
	}
	if err = s.store.CreateDepartment(d); err != nil {
		return d, fmt.Errorf("create department: %w", err)
	}
	return d, nil
}
func (s *Service) UpdateDetails(id string, details domain.DepartmentDetails) (domain.Department, error) {
	d, err := s.store.GetDepartment(id)
	if err != nil {
		return d, err
	}
	d, err = d.UpdateDetails(details, s.now())
	if err != nil {
		return d, err
	}
	return d, s.store.PutDepartment(d)
}
func (s *Service) Activate(id string) (domain.Department, error) {
	d, err := s.store.GetDepartment(id)
	if err != nil {
		return d, err
	}
	d, err = d.Activate(s.now())
	if err != nil {
		return d, fmt.Errorf("activate department: %w", err)
	}
	return d, s.store.PutDepartment(d)
}
func (s *Service) Deactivate(id string) (domain.Department, error) {
	accounts, err := s.store.ListAccounts(id, "", domain.AccountActive)
	if err != nil {
		return domain.Department{}, err
	}
	if len(accounts) > 0 {
		return domain.Department{}, fmt.Errorf("%w: active accounts remain", domain.ErrDepartmentState)
	}
	d, err := s.store.GetDepartment(id)
	if err != nil {
		return d, err
	}
	d, err = d.Deactivate(s.now())
	if err != nil {
		return d, err
	}
	return d, s.store.PutDepartment(d)
}
func (s *Service) Get(id string) (domain.Department, error) { return s.store.GetDepartment(id) }
func (s *Service) List(status domain.DepartmentStatus) ([]domain.Department, error) {
	return s.store.ListDepartments(status)
}
