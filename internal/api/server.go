package api

import (
	"encoding/json"
	portal "hospitalportal"
	"hospitalportal/internal/account"
	"hospitalportal/internal/domain"
	"hospitalportal/internal/schedule"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	portal *portal.Portal
	logger *slog.Logger
	mux    *http.ServeMux
}

func New(p *portal.Portal, logger *slog.Logger) *Server {
	s := &Server{portal: p, logger: logger, mux: http.NewServeMux()}
	s.routes()
	return s
}
func (s *Server) Handler() http.Handler { return s.logging(s.mux) }
func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.health)
	s.mux.HandleFunc("GET /api/openapi.yaml", s.openAPI)
	s.mux.HandleFunc("GET /", s.index)
	s.mux.HandleFunc("GET /api/accounts", s.listAccounts)
	s.mux.HandleFunc("POST /api/accounts", s.createAccount)
	s.mux.HandleFunc("GET /api/accounts/{id}", s.getAccount)
	s.mux.HandleFunc("POST /api/accounts/{id}/activate", s.activateAccount)
	s.mux.HandleFunc("POST /api/accounts/{id}/suspend", s.suspendAccount)
	s.mux.HandleFunc("GET /api/departments", s.listDepartments)
	s.mux.HandleFunc("POST /api/departments", s.createDepartment)
	s.mux.HandleFunc("GET /api/departments/{id}", s.getDepartment)
	s.mux.HandleFunc("PUT /api/departments/{id}", s.updateDepartment)
	s.mux.HandleFunc("POST /api/departments/{id}/activate", s.activateDepartment)
	s.mux.HandleFunc("GET /api/shifts", s.listShifts)
	s.mux.HandleFunc("POST /api/shifts", s.createShift)
	s.mux.HandleFunc("POST /api/shifts/{id}/publish", s.publishShift)
}
func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Info("request completed", "request_id", r.Header.Get("X-Request-ID"), "method", r.Method, "path", r.URL.Path, "duration_ms", time.Since(started).Milliseconds())
	})
}
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
func (s *Server) openAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	_, _ = io.WriteString(w, OpenAPISpec)
}
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = io.WriteString(w, PortalHTML)
}
func (s *Server) listAccounts(w http.ResponseWriter, r *http.Request) {
	values, err := s.portal.Accounts.List(r.URL.Query().Get("department_id"), domain.Role(r.URL.Query().Get("role")), domain.AccountStatus(r.URL.Query().Get("status")))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": values, "count": len(values)})
}
func (s *Server) getAccount(w http.ResponseWriter, r *http.Request) {
	value, err := s.portal.Accounts.Get(r.PathValue("id"))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func decode[T any](r *http.Request) (T, error) {
	var value T
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&value)
	return value, err
}
func (s *Server) createAccount(w http.ResponseWriter, r *http.Request) {
	body, err := decode[struct {
		ID             string      `json:"id"`
		DisplayName    string      `json:"display_name"`
		EmployeeNumber string      `json:"employee_number"`
		Email          string      `json:"email"`
		DepartmentID   string      `json:"department_id"`
		Role           domain.Role `json:"role"`
	}](r)
	if err != nil {
		writeError(w, r, s.logger, domain.ErrInvalidAccount)
		return
	}
	value, err := s.portal.Accounts.Create(accountCommand(body.ID, body.DisplayName, body.EmployeeNumber, body.Email, body.DepartmentID, body.Role))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusCreated, value)
}
func accountCommand(id, name, employee, email, department string, role domain.Role) account.CreateCommand {
	return account.CreateCommand{ID: id, DisplayName: name, EmployeeNumber: employee, Email: email, DepartmentID: department, Role: role}
}
func (s *Server) activateAccount(w http.ResponseWriter, r *http.Request) {
	value, err := s.portal.Accounts.Activate(r.PathValue("id"))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func (s *Server) suspendAccount(w http.ResponseWriter, r *http.Request) {
	value, err := s.portal.Accounts.Suspend(r.PathValue("id"))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func (s *Server) listDepartments(w http.ResponseWriter, r *http.Request) {
	values, err := s.portal.Departments.List(domain.DepartmentStatus(r.URL.Query().Get("status")))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": values, "count": len(values)})
}
func (s *Server) getDepartment(w http.ResponseWriter, r *http.Request) {
	value, err := s.portal.Departments.Get(r.PathValue("id"))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func (s *Server) createDepartment(w http.ResponseWriter, r *http.Request) {
	body, err := decode[struct {
		ID   string `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}](r)
	if err != nil {
		writeError(w, r, s.logger, domain.ErrInvalidDepartment)
		return
	}
	value, err := s.portal.Departments.Create(body.ID, body.Code, body.Name)
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusCreated, value)
}
func (s *Server) updateDepartment(w http.ResponseWriter, r *http.Request) {
	body, err := decode[struct {
		Description   string   `json:"description"`
		Location      string   `json:"location"`
		Phone         string   `json:"phone"`
		Email         string   `json:"email"`
		HeadAccountID string   `json:"head_account_id"`
		Services      []string `json:"services"`
	}](r)
	if err != nil {
		writeError(w, r, s.logger, domain.ErrInvalidDepartment)
		return
	}
	value, err := s.portal.Departments.UpdateDetails(r.PathValue("id"), domain.DepartmentDetails{Description: body.Description, Location: body.Location, Phone: body.Phone, Email: body.Email, HeadAccountID: body.HeadAccountID, Services: body.Services})
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func (s *Server) activateDepartment(w http.ResponseWriter, r *http.Request) {
	value, err := s.portal.Departments.Activate(r.PathValue("id"))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func (s *Server) listShifts(w http.ResponseWriter, r *http.Request) {
	from, _ := time.Parse(time.RFC3339, r.URL.Query().Get("from"))
	to, _ := time.Parse(time.RFC3339, r.URL.Query().Get("to"))
	values, err := s.portal.Schedules.List(r.URL.Query().Get("department_id"), r.URL.Query().Get("account_id"), from, to)
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": values, "count": len(values)})
}
func (s *Server) createShift(w http.ResponseWriter, r *http.Request) {
	body, err := decode[struct {
		ID           string    `json:"id"`
		DepartmentID string    `json:"department_id"`
		AccountID    string    `json:"account_id"`
		Title        string    `json:"title"`
		StartAt      time.Time `json:"start_at"`
		EndAt        time.Time `json:"end_at"`
	}](r)
	if err != nil {
		writeError(w, r, s.logger, domain.ErrInvalidShift)
		return
	}
	value, err := s.portal.Schedules.Create(schedule.CreateCommand{ID: body.ID, DepartmentID: body.DepartmentID, AccountID: body.AccountID, Title: body.Title, StartAt: body.StartAt, EndAt: body.EndAt})
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusCreated, value)
}
func (s *Server) publishShift(w http.ResponseWriter, r *http.Request) {
	value, err := s.portal.Schedules.Publish(r.PathValue("id"))
	if err != nil {
		writeError(w, r, s.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, value)
}
func normalizeQuery(value string) string { return strings.TrimSpace(strings.ToLower(value)) }
