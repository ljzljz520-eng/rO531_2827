package api

import (
	"encoding/json"
	"errors"
	"hospitalportal/internal/domain"
	"log/slog"
	"net/http"
)

type ErrorBody struct {
	Error APIError `json:"error"`
}
type APIError struct {
	Code      string            `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id,omitempty"`
	Details   map[string]string `json:"details,omitempty"`
}

func writeError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	status, code := classifyError(err)
	requestID := r.Header.Get("X-Request-ID")
	logger.Error("request failed", "request_id", requestID, "method", r.Method, "path", r.URL.Path, "status", status, "error_code", code, "error", err)
	writeJSON(w, status, ErrorBody{Error: APIError{Code: code, Message: http.StatusText(status), RequestID: requestID}})
}
func classifyError(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrAccountNotFound), errors.Is(err, domain.ErrDepartmentNotFound), errors.Is(err, domain.ErrShiftNotFound):
		return http.StatusNotFound, "resource_not_found"
	case errors.Is(err, domain.ErrPermissionDenied):
		return http.StatusForbidden, "permission_denied"
	case errors.Is(err, domain.ErrDuplicateAccount):
		return http.StatusConflict, "duplicate_account"
	case errors.Is(err, domain.ErrAccountState), errors.Is(err, domain.ErrDepartmentState), errors.Is(err, domain.ErrShiftConflict):
		return http.StatusConflict, "state_conflict"
	case errors.Is(err, domain.ErrInvalidAccount), errors.Is(err, domain.ErrInvalidDepartment), errors.Is(err, domain.ErrInvalidShift):
		return http.StatusUnprocessableEntity, "validation_failed"
	default:
		return http.StatusInternalServerError, "internal_error"
	}
}
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
