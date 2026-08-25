package account

import (
	"hospitalportal/internal/domain"
	"testing"
	"time"
)

func TestAssignmentPolicyRejectsDuplicateEmployee(t *testing.T) {
	a, err := domain.NewUserAccount("a", "A", "E1", "a@h.test", domain.RoleDoctor, "d", time.Unix(1, 0))
	if err != nil {
		t.Fatal(err)
	}
	b, err := domain.NewUserAccount("b", "B", "E1", "b@h.test", domain.RoleNurse, "d", time.Unix(1, 0))
	if err != nil {
		t.Fatal(err)
	}
	if err = DefaultPolicy().ValidateCandidate(b, []domain.UserAccount{a}); err == nil {
		t.Fatal("expected duplicate rejection")
	}
}
