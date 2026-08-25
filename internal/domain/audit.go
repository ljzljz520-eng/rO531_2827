package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInvalidAuditRecord = errors.New("invalid audit record")

type AuditRecord struct {
	ID          string            `json:"id"`
	ActorID     string            `json:"actor_id"`
	Action      string            `json:"action"`
	SubjectType string            `json:"subject_type"`
	SubjectID   string            `json:"subject_id"`
	Outcome     string            `json:"outcome"`
	Fields      map[string]string `json:"fields"`
	OccurredAt  time.Time         `json:"occurred_at"`
}

func NewAuditRecord(id, actor, action, subjectType, subjectID, outcome string, fields map[string]string, now time.Time) (AuditRecord, error) {
	r := AuditRecord{ID: strings.TrimSpace(id), ActorID: strings.TrimSpace(actor), Action: strings.TrimSpace(action), SubjectType: strings.TrimSpace(subjectType), SubjectID: strings.TrimSpace(subjectID), Outcome: strings.TrimSpace(outcome), Fields: copyFields(fields), OccurredAt: now.UTC()}
	if err := r.Validate(); err != nil {
		return AuditRecord{}, err
	}
	return r, nil
}
func (r AuditRecord) Validate() error {
	if r.ID == "" || r.ActorID == "" || r.Action == "" || r.SubjectType == "" || r.SubjectID == "" {
		return fmt.Errorf("%w: identity fields", ErrInvalidAuditRecord)
	}
	if r.Outcome != "success" && r.Outcome != "failure" {
		return fmt.Errorf("%w: outcome", ErrInvalidAuditRecord)
	}
	for k := range r.Fields {
		if strings.TrimSpace(k) == "" {
			return fmt.Errorf("%w: empty field name", ErrInvalidAuditRecord)
		}
	}
	return nil
}
func copyFields(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
