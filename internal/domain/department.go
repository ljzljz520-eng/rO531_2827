package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrDepartmentNotFound = errors.New("department not found")
	ErrInvalidDepartment  = errors.New("invalid department")
	ErrDepartmentState    = errors.New("department state conflict")
)

type DepartmentStatus string

const (
	DepartmentDraft    DepartmentStatus = "draft"
	DepartmentActive   DepartmentStatus = "active"
	DepartmentInactive DepartmentStatus = "inactive"
)

type Department struct {
	ID            string           `json:"id"`
	Code          string           `json:"code"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Location      string           `json:"location"`
	Phone         string           `json:"phone"`
	Email         string           `json:"email"`
	HeadAccountID string           `json:"head_account_id"`
	Status        DepartmentStatus `json:"status"`
	Services      []string         `json:"services"`
	Version       int              `json:"version"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

func NewDepartment(id, code, name string, now time.Time) (Department, error) {
	d := Department{ID: strings.TrimSpace(id), Code: strings.ToUpper(strings.TrimSpace(code)), Name: strings.TrimSpace(name), Status: DepartmentDraft, Version: 1, CreatedAt: now.UTC(), UpdatedAt: now.UTC()}
	if err := d.Validate(); err != nil {
		return Department{}, err
	}
	return d, nil
}
func (d Department) Validate() error {
	if d.ID == "" || d.Code == "" || d.Name == "" {
		return fmt.Errorf("%w: identity fields required", ErrInvalidDepartment)
	}
	if len(d.Code) < 2 || len(d.Code) > 12 {
		return fmt.Errorf("%w: code length", ErrInvalidDepartment)
	}
	switch d.Status {
	case DepartmentDraft, DepartmentActive, DepartmentInactive:
	default:
		return fmt.Errorf("%w: status", ErrInvalidDepartment)
	}
	return nil
}

type DepartmentDetails struct {
	Description, Location, Phone, Email, HeadAccountID string
	Services                                           []string
}

func (d Department) UpdateDetails(v DepartmentDetails, now time.Time) (Department, error) {
	if strings.TrimSpace(v.Location) == "" || strings.TrimSpace(v.Phone) == "" {
		return d, fmt.Errorf("%w: contact details required", ErrInvalidDepartment)
	}
	if v.Email != "" && !strings.Contains(v.Email, "@") {
		return d, fmt.Errorf("%w: email", ErrInvalidDepartment)
	}
	d.Description = strings.TrimSpace(v.Description)
	d.Location = strings.TrimSpace(v.Location)
	d.Phone = strings.TrimSpace(v.Phone)
	d.Email = strings.TrimSpace(v.Email)
	d.HeadAccountID = strings.TrimSpace(v.HeadAccountID)
	d.Services = normalizeServices(v.Services)
	d.Version++
	d.UpdatedAt = now.UTC()
	return d, nil
}
func normalizeServices(values []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" && !seen[value] {
			seen[value] = true
			out = append(out, value)
		}
	}
	return out
}
func (d Department) Activate(now time.Time) (Department, error) {
	if d.Status != DepartmentDraft && d.Status != DepartmentInactive {
		return d, fmt.Errorf("%w: cannot activate", ErrDepartmentState)
	}
	if d.Location == "" || d.Phone == "" {
		return d, fmt.Errorf("%w: incomplete details", ErrDepartmentState)
	}
	d.Status = DepartmentActive
	d.Version++
	d.UpdatedAt = now.UTC()
	return d, nil
}
func (d Department) Deactivate(now time.Time) (Department, error) {
	if d.Status != DepartmentActive {
		return d, fmt.Errorf("%w: cannot deactivate", ErrDepartmentState)
	}
	d.Status = DepartmentInactive
	d.Version++
	d.UpdatedAt = now.UTC()
	return d, nil
}
