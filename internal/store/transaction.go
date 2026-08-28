package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"hospitalportal/internal/domain"
)

type UnitOfWork struct{ tx *bbolt.Tx }

func (s *Store) Transact(fn func(*UnitOfWork) error) error {
	return s.Update(func(tx *bbolt.Tx) error { return fn(&UnitOfWork{tx: tx}) })
}
func (u *UnitOfWork) Account(id string) (domain.UserAccount, error) {
	var v domain.UserAccount
	err := getJSON(u.tx, accountsBucket, id, &v, domain.ErrAccountNotFound)
	return v, err
}
func (u *UnitOfWork) Department(id string) (domain.Department, error) {
	var v domain.Department
	err := getJSON(u.tx, departmentsBucket, id, &v, domain.ErrDepartmentNotFound)
	return v, err
}
func (u *UnitOfWork) Shift(id string) (domain.DutyShift, error) {
	var v domain.DutyShift
	err := getJSON(u.tx, shiftsBucket, id, &v, domain.ErrShiftNotFound)
	return v, err
}
func (u *UnitOfWork) PutAccount(v domain.UserAccount) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return putJSON(u.tx, accountsBucket, v.ID, v)
}
func (u *UnitOfWork) PutDepartment(v domain.Department) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return putJSON(u.tx, departmentsBucket, v.ID, v)
}
func (u *UnitOfWork) PutShift(v domain.DutyShift) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return putJSON(u.tx, shiftsBucket, v.ID, v)
}
func (u *UnitOfWork) AppendAudit(v domain.AuditRecord) error {
	if err := v.Validate(); err != nil {
		return err
	}
	if u.tx.Bucket(auditsBucket).Get([]byte(v.ID)) != nil {
		return fmt.Errorf("audit id exists")
	}
	return putJSON(u.tx, auditsBucket, v.ID, v)
}
