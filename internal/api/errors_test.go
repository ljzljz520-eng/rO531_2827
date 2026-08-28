package api

import (
	"errors"
	"hospitalportal/internal/domain"
	"net/http"
	"testing"
)

func TestClassifyWrappedStateError(t *testing.T) {
	status, code := classifyError(errors.Join(errors.New("service"), domain.ErrAccountState))
	if status != http.StatusConflict || code != "state_conflict" {
		t.Fatalf("got %d %s", status, code)
	}
}
