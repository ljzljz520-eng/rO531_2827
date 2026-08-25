package audit

import (
	"fmt"
	"hospitalportal/internal/domain"
	"hospitalportal/internal/store"
	"sync"
	"time"
)

type Service struct {
	store    *store.Store
	now      func() time.Time
	mu       sync.Mutex
	sequence uint64
}

func New(s *store.Store, now func() time.Time) *Service { return &Service{store: s, now: now} }
func (s *Service) Record(actor, action, subjectType, subjectID, outcome string, fields map[string]string) error {
	s.mu.Lock()
	s.sequence++
	id := fmt.Sprintf("audit-%06d", s.sequence)
	s.mu.Unlock()
	record, err := domain.NewAuditRecord(id, actor, action, subjectType, subjectID, outcome, fields, s.now())
	if err != nil {
		return err
	}
	return s.store.AppendAudit(record)
}
func (s *Service) ForSubject(subjectType, subjectID string) ([]domain.AuditRecord, error) {
	return s.store.ListAudits(subjectType, subjectID)
}
