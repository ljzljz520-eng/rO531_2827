package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Role string

const (
	RoleDoctor        Role = "doctor"
	RoleNurse         Role = "nurse"
	RoleAdministrator Role = "administrator"
)

type AccountStatus string

const (
	AccountPending   AccountStatus = "pending"
	AccountActive    AccountStatus = "active"
	AccountSuspended AccountStatus = "suspended"
)

var (
	ErrInvalidAccount   = errors.New("invalid account")
	ErrAccountNotFound  = errors.New("account not found")
	ErrAccountState     = errors.New("account state conflict")
	ErrPermissionDenied = errors.New("permission denied")
	ErrDuplicateAccount = errors.New("duplicate account")
)

type UserAccount struct {
	ID             string        `json:"id"`
	DisplayName    string        `json:"display_name"`
	EmployeeNumber string        `json:"employee_number"`
	Email          string        `json:"email"`
	Phone          string        `json:"phone"`
	Role           Role          `json:"role"`
	DepartmentID   string        `json:"department_id"`
	Status         AccountStatus `json:"status"`
	Version        int           `json:"version"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type AccountProfile struct {
	DisplayName  string `json:"display_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	DepartmentID string `json:"department_id"`
}

func NewUserAccount(id, name, employee, email string, role Role, departmentID string, now time.Time) (UserAccount, error) {
	a := UserAccount{ID: strings.TrimSpace(id), DisplayName: strings.TrimSpace(name), EmployeeNumber: strings.TrimSpace(employee), Email: strings.TrimSpace(email), Role: role, DepartmentID: strings.TrimSpace(departmentID), Status: AccountPending, Version: 1, CreatedAt: now.UTC(), UpdatedAt: now.UTC()}
	if err := a.Validate(); err != nil {
		return UserAccount{}, err
	}
	return a, nil
}

func (a UserAccount) Validate() error {
	if a.ID == "" || a.DisplayName == "" || a.EmployeeNumber == "" {
		return fmt.Errorf("%w: identity fields are required", ErrInvalidAccount)
	}
	if !strings.Contains(a.Email, "@") {
		return fmt.Errorf("%w: email format", ErrInvalidAccount)
	}
	switch a.Role {
	case RoleDoctor, RoleNurse, RoleAdministrator:
	default:
		return fmt.Errorf("%w: unsupported role", ErrInvalidAccount)
	}
	if a.Role != RoleAdministrator && a.DepartmentID == "" {
		return fmt.Errorf("%w: clinical account requires department", ErrInvalidAccount)
	}
	switch a.Status {
	case AccountPending, AccountActive, AccountSuspended:
	default:
		return fmt.Errorf("%w: unsupported status", ErrInvalidAccount)
	}
	return nil
}

func (a UserAccount) Activate(now time.Time) (UserAccount, error) {
	if a.Status != AccountPending && a.Status != AccountSuspended {
		return a, fmt.Errorf("%w: cannot activate %s", ErrAccountState, a.Status)
	}
	a.Status = AccountActive
	a.Version++
	a.UpdatedAt = now.UTC()
	return a, nil
}
func (a UserAccount) Suspend(now time.Time) (UserAccount, error) {
	if a.Status != AccountActive {
		return a, fmt.Errorf("%w: cannot suspend %s", ErrAccountState, a.Status)
	}
	a.Status = AccountSuspended
	a.Version++
	a.UpdatedAt = now.UTC()
	return a, nil
}
func (a UserAccount) UpdateProfile(p AccountProfile, now time.Time) (UserAccount, error) {
	if strings.TrimSpace(p.DisplayName) == "" || !strings.Contains(p.Email, "@") {
		return a, fmt.Errorf("%w: invalid profile", ErrInvalidAccount)
	}
	a.DisplayName = strings.TrimSpace(p.DisplayName)
	a.Email = strings.TrimSpace(p.Email)
	a.Phone = strings.TrimSpace(p.Phone)
	if a.Role != RoleAdministrator && strings.TrimSpace(p.DepartmentID) == "" {
		return a, fmt.Errorf("%w: department required", ErrInvalidAccount)
	}
	a.DepartmentID = strings.TrimSpace(p.DepartmentID)
	a.Version++
	a.UpdatedAt = now.UTC()
	return a, nil
}
func (a UserAccount) Can(action string) bool {
	if a.Status != AccountActive {
		return false
	}
	if a.Role == RoleAdministrator {
		return true
	}
	switch action {
	case "department.read", "schedule.read":
		return true
	case "schedule.manage":
		return a.Role == RoleNurse
	case "clinical.approve":
		return a.Role == RoleDoctor
	}
	return false
}
