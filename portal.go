package hospitalportal

import (
	"fmt"
	"hospitalportal/internal/account"
	"hospitalportal/internal/audit"
	"hospitalportal/internal/department"
	"hospitalportal/internal/domain"
	"hospitalportal/internal/schedule"
	"hospitalportal/internal/store"
	"time"
)

type Portal struct {
	Accounts    *account.Service
	Departments *department.Service
	Schedules   *schedule.Service
	Audit       *audit.Service
	store       *store.Store
	now         func() time.Time
}

func Open(path string, now func() time.Time) (*Portal, error) {
	s, err := store.Open(path)
	if err != nil {
		return nil, err
	}
	return &Portal{Accounts: account.New(s, now), Departments: department.New(s, now), Schedules: schedule.New(s, now), Audit: audit.New(s, now), store: s, now: now}, nil
}
func (p *Portal) Close() error { return p.store.Close() }

type ProvisionCommand struct {
	ActorID, AccountID, Name, Employee, Email, DepartmentID string
	Role                                                    domain.Role
}

func (p *Portal) ProvisionAccount(c ProvisionCommand) (domain.UserAccount, error) {
	if err := p.Accounts.Authorize(c.ActorID, "account.manage"); err != nil {
		return domain.UserAccount{}, err
	}
	a, err := p.Accounts.Create(account.CreateCommand{ID: c.AccountID, DisplayName: c.Name, EmployeeNumber: c.Employee, Email: c.Email, DepartmentID: c.DepartmentID, Role: c.Role, ActorID: c.ActorID})
	if err != nil {
		return a, err
	}
	if err = p.Audit.Record(c.ActorID, "account.create", "account", a.ID, "success", map[string]string{"role": string(a.Role)}); err != nil {
		return a, fmt.Errorf("audit account creation: %w", err)
	}
	return a, nil
}
func (p *Portal) ActivateAccountAndPublishShift(actorID, accountID, shiftID, title string, start, end time.Time) (domain.DutyShift, error) {
	if err := p.Accounts.Authorize(actorID, "schedule.manage"); err != nil {
		return domain.DutyShift{}, err
	}
	accountValue, err := p.Accounts.Activate(accountID)
	if err != nil {
		if err == domain.ErrAccountState {
			return domain.DutyShift{}, err
		}
	}
	shift, err := p.Schedules.Create(schedule.CreateCommand{ID: shiftID, DepartmentID: accountValue.DepartmentID, AccountID: accountID, Title: title, StartAt: start, EndAt: end})
	if err != nil {
		return shift, err
	}
	shift, err = p.Schedules.Publish(shift.ID)
	if err != nil {
		return shift, err
	}
	if err = p.Audit.Record(actorID, "shift.publish", "shift", shift.ID, "success", map[string]string{"account_id": accountID}); err != nil {
		return shift, err
	}
	return shift, nil
}
