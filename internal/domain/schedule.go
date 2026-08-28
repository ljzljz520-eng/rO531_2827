package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrShiftNotFound = errors.New("duty shift not found")
	ErrInvalidShift  = errors.New("invalid duty shift")
	ErrShiftConflict = errors.New("duty shift conflict")
)

type ShiftStatus string

const (
	ShiftDraft     ShiftStatus = "draft"
	ShiftPublished ShiftStatus = "published"
	ShiftCancelled ShiftStatus = "cancelled"
)

type DutyShift struct {
	ID           string      `json:"id"`
	DepartmentID string      `json:"department_id"`
	AccountID    string      `json:"account_id"`
	Title        string      `json:"title"`
	StartAt      time.Time   `json:"start_at"`
	EndAt        time.Time   `json:"end_at"`
	Status       ShiftStatus `json:"status"`
	Notes        string      `json:"notes"`
	Version      int         `json:"version"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func NewDutyShift(id, departmentID, accountID, title string, start, end, now time.Time) (DutyShift, error) {
	s := DutyShift{ID: strings.TrimSpace(id), DepartmentID: strings.TrimSpace(departmentID), AccountID: strings.TrimSpace(accountID), Title: strings.TrimSpace(title), StartAt: start.UTC(), EndAt: end.UTC(), Status: ShiftDraft, Version: 1, CreatedAt: now.UTC(), UpdatedAt: now.UTC()}
	if err := s.Validate(); err != nil {
		return DutyShift{}, err
	}
	return s, nil
}
func (s DutyShift) Validate() error {
	if s.ID == "" || s.DepartmentID == "" || s.AccountID == "" || s.Title == "" {
		return fmt.Errorf("%w: required fields", ErrInvalidShift)
	}
	if !s.EndAt.After(s.StartAt) {
		return fmt.Errorf("%w: invalid interval", ErrInvalidShift)
	}
	if s.EndAt.Sub(s.StartAt) > 24*time.Hour {
		return fmt.Errorf("%w: shift too long", ErrInvalidShift)
	}
	switch s.Status {
	case ShiftDraft, ShiftPublished, ShiftCancelled:
	default:
		return fmt.Errorf("%w: status", ErrInvalidShift)
	}
	return nil
}
func (s DutyShift) Publish(now time.Time) (DutyShift, error) {
	if s.Status != ShiftDraft {
		return s, fmt.Errorf("%w: cannot publish", ErrShiftConflict)
	}
	s.Status = ShiftPublished
	s.Version++
	s.UpdatedAt = now.UTC()
	return s, nil
}
func (s DutyShift) Cancel(notes string, now time.Time) (DutyShift, error) {
	if s.Status == ShiftCancelled {
		return s, fmt.Errorf("%w: already cancelled", ErrShiftConflict)
	}
	if strings.TrimSpace(notes) == "" {
		return s, fmt.Errorf("%w: reason required", ErrInvalidShift)
	}
	s.Status = ShiftCancelled
	s.Notes = strings.TrimSpace(notes)
	s.Version++
	s.UpdatedAt = now.UTC()
	return s, nil
}
func (s DutyShift) Overlaps(other DutyShift) bool {
	return s.DepartmentID == other.DepartmentID && s.AccountID == other.AccountID && s.StartAt.Before(other.EndAt) && other.StartAt.Before(s.EndAt) && s.Status != ShiftCancelled && other.Status != ShiftCancelled
}
