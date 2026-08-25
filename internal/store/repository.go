package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"hospitalportal/internal/domain"
	"time"
)

func (s *Store) CreateAccount(v domain.UserAccount) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(accountsBucket)
		if b.Get([]byte(v.ID)) != nil {
			return domain.ErrDuplicateAccount
		}
		return putJSON(tx, accountsBucket, v.ID, v)
	})
}
func (s *Store) PutAccount(v domain.UserAccount) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error { return putJSON(tx, accountsBucket, v.ID, v) })
}
func (s *Store) GetAccount(id string) (domain.UserAccount, error) {
	var v domain.UserAccount
	err := s.View(func(tx *bbolt.Tx) error { return getJSON(tx, accountsBucket, id, &v, domain.ErrAccountNotFound) })
	return v, err
}
func (s *Store) FindAccountByEmployee(employee string) (domain.UserAccount, error) {
	var found domain.UserAccount
	err := s.View(func(tx *bbolt.Tx) error {
		values, err := listJSON[domain.UserAccount](tx, accountsBucket, nil)
		if err != nil {
			return err
		}
		for _, value := range values {
			if value.EmployeeNumber == employee {
				found = value
				return nil
			}
		}
		return domain.ErrAccountNotFound
	})
	return found, err
}
func (s *Store) ListAccounts(departmentID string, role domain.Role, status domain.AccountStatus) ([]domain.UserAccount, error) {
	var out []domain.UserAccount
	err := s.View(func(tx *bbolt.Tx) error {
		values, err := listJSON[domain.UserAccount](tx, accountsBucket, func(v domain.UserAccount) bool {
			if departmentID != "" && v.DepartmentID != departmentID {
				return false
			}
			if role != "" && v.Role != role {
				return false
			}
			if status != "" && v.Status != status {
				return false
			}
			return true
		})
		out = values
		return err
	})
	sortAccounts(out)
	return out, err
}
func (s *Store) CreateDepartment(v domain.Department) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(departmentsBucket)
		if b.Get([]byte(v.ID)) != nil {
			return fmt.Errorf("department exists")
		}
		return putJSON(tx, departmentsBucket, v.ID, v)
	})
}
func (s *Store) PutDepartment(v domain.Department) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error { return putJSON(tx, departmentsBucket, v.ID, v) })
}
func (s *Store) GetDepartment(id string) (domain.Department, error) {
	var v domain.Department
	err := s.View(func(tx *bbolt.Tx) error { return getJSON(tx, departmentsBucket, id, &v, domain.ErrDepartmentNotFound) })
	return v, err
}
func (s *Store) ListDepartments(status domain.DepartmentStatus) ([]domain.Department, error) {
	var out []domain.Department
	err := s.View(func(tx *bbolt.Tx) error {
		values, err := listJSON[domain.Department](tx, departmentsBucket, func(v domain.Department) bool { return status == "" || v.Status == status })
		out = values
		return err
	})
	sortDepartments(out)
	return out, err
}
func (s *Store) CreateShift(v domain.DutyShift) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error {
		if tx.Bucket(shiftsBucket).Get([]byte(v.ID)) != nil {
			return fmt.Errorf("shift exists")
		}
		return putJSON(tx, shiftsBucket, v.ID, v)
	})
}
func (s *Store) PutShift(v domain.DutyShift) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error { return putJSON(tx, shiftsBucket, v.ID, v) })
}
func (s *Store) GetShift(id string) (domain.DutyShift, error) {
	var v domain.DutyShift
	err := s.View(func(tx *bbolt.Tx) error { return getJSON(tx, shiftsBucket, id, &v, domain.ErrShiftNotFound) })
	return v, err
}
func (s *Store) ListShifts(departmentID, accountID string, from, to time.Time) ([]domain.DutyShift, error) {
	var out []domain.DutyShift
	err := s.View(func(tx *bbolt.Tx) error {
		values, err := listJSON[domain.DutyShift](tx, shiftsBucket, func(v domain.DutyShift) bool {
			if departmentID != "" && v.DepartmentID != departmentID {
				return false
			}
			if accountID != "" && v.AccountID != accountID {
				return false
			}
			if !from.IsZero() && v.EndAt.Before(from) {
				return false
			}
			if !to.IsZero() && v.StartAt.After(to) {
				return false
			}
			return true
		})
		out = values
		return err
	})
	sortShifts(out)
	return out, err
}
func (s *Store) AppendAudit(v domain.AuditRecord) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return s.Update(func(tx *bbolt.Tx) error { return putJSON(tx, auditsBucket, v.ID, v) })
}
func (s *Store) ListAudits(subjectType, subjectID string) ([]domain.AuditRecord, error) {
	var out []domain.AuditRecord
	err := s.View(func(tx *bbolt.Tx) error {
		values, err := listJSON[domain.AuditRecord](tx, auditsBucket, func(v domain.AuditRecord) bool {
			return (subjectType == "" || v.SubjectType == subjectType) && (subjectID == "" || v.SubjectID == subjectID)
		})
		out = values
		return err
	})
	sortAudits(out)
	return out, err
}
