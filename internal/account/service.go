package account

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
	ID, DisplayName, EmployeeNumber, Email, DepartmentID, ActorID string
	Role                                                          domain.Role
}

func (s *Service) Create(c CreateCommand) (domain.UserAccount, error) {
	if _, err := s.store.FindAccountByEmployee(c.EmployeeNumber); err == nil {
		return domain.UserAccount{}, domain.ErrDuplicateAccount
	}
	a, err := domain.NewUserAccount(c.ID, c.DisplayName, c.EmployeeNumber, c.Email, c.Role, c.DepartmentID, s.now())
	if err != nil {
		return domain.UserAccount{}, err
	}
	if err = s.store.CreateAccount(a); err != nil {
		return domain.UserAccount{}, fmt.Errorf("create account: %w", err)
	}
	return a, nil
}
func (s *Service) Activate(id string) (domain.UserAccount, error) {
	a, err := s.store.GetAccount(id)
	if err != nil {
		return domain.UserAccount{}, fmt.Errorf("load account: %w", err)
	}
	a, err = a.Activate(s.now())
	if err != nil {
		return domain.UserAccount{}, fmt.Errorf("activate account: %w", err)
	}
	if err = s.store.PutAccount(a); err != nil {
		return domain.UserAccount{}, fmt.Errorf("save account: %w", err)
	}
	return a, nil
}
func (s *Service) Suspend(id string) (domain.UserAccount, error) {
	a, err := s.store.GetAccount(id)
	if err != nil {
		return domain.UserAccount{}, err
	}
	a, err = a.Suspend(s.now())
	if err != nil {
		return domain.UserAccount{}, fmt.Errorf("suspend account: %w", err)
	}
	return a, s.store.PutAccount(a)
}
func (s *Service) Update(id string, p domain.AccountProfile) (domain.UserAccount, error) {
	a, err := s.store.GetAccount(id)
	if err != nil {
		return domain.UserAccount{}, err
	}
	a, err = a.UpdateProfile(p, s.now())
	if err != nil {
		return domain.UserAccount{}, err
	}
	return a, s.store.PutAccount(a)
}
func (s *Service) Get(id string) (domain.UserAccount, error) { return s.store.GetAccount(id) }
func (s *Service) List(department string, role domain.Role, status domain.AccountStatus) ([]domain.UserAccount, error) {
	return s.store.ListAccounts(department, role, status)
}
func (s *Service) Authorize(id, action string) error {
	a, err := s.store.GetAccount(id)
	if err != nil {
		return err
	}
	if !a.Can(action) {
		return fmt.Errorf("%w: %s cannot %s", domain.ErrPermissionDenied, id, action)
	}
	return nil
}
