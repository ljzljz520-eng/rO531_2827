package account

import (
	"fmt"
	"hospitalportal/internal/domain"
)

type AssignmentPolicy struct {
	MaximumAdministrators     int
	RequireClinicalDepartment bool
}

func DefaultPolicy() AssignmentPolicy {
	return AssignmentPolicy{MaximumAdministrators: 12, RequireClinicalDepartment: true}
}
func (p AssignmentPolicy) ValidateCandidate(candidate domain.UserAccount, existing []domain.UserAccount) error {
	if p.RequireClinicalDepartment && candidate.Role != domain.RoleAdministrator && candidate.DepartmentID == "" {
		return fmt.Errorf("%w: clinical department required", domain.ErrInvalidAccount)
	}
	admins := 0
	for _, a := range existing {
		if a.Role == domain.RoleAdministrator && a.Status != domain.AccountSuspended {
			admins++
		}
		if a.EmployeeNumber == candidate.EmployeeNumber && a.ID != candidate.ID {
			return domain.ErrDuplicateAccount
		}
	}
	if candidate.Role == domain.RoleAdministrator && admins >= p.MaximumAdministrators {
		return fmt.Errorf("%w: administrator limit", domain.ErrInvalidAccount)
	}
	return nil
}
func VisibleTo(actor, target domain.UserAccount) bool {
	if actor.Status != domain.AccountActive {
		return false
	}
	if actor.Role == domain.RoleAdministrator {
		return true
	}
	if actor.DepartmentID != target.DepartmentID {
		return false
	}
	return target.Role == domain.RoleDoctor || target.Role == domain.RoleNurse
}
