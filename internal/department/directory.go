package department

import (
	"hospitalportal/internal/domain"
	"sort"
	"strings"
)

type DirectoryEntry struct {
	ID, Code, Name, Location, Phone, HeadAccountID string
	Services                                       []string
	ActiveAccounts                                 int
}

func BuildDirectory(departments []domain.Department, accounts []domain.UserAccount, query string) []DirectoryEntry {
	counts := map[string]int{}
	for _, a := range accounts {
		if a.Status == domain.AccountActive {
			counts[a.DepartmentID]++
		}
	}
	query = strings.ToLower(strings.TrimSpace(query))
	out := []DirectoryEntry{}
	for _, d := range departments {
		if d.Status != domain.DepartmentActive {
			continue
		}
		if query != "" && !strings.Contains(strings.ToLower(d.Name+" "+d.Code+" "+strings.Join(d.Services, " ")), query) {
			continue
		}
		out = append(out, DirectoryEntry{ID: d.ID, Code: d.Code, Name: d.Name, Location: d.Location, Phone: d.Phone, HeadAccountID: d.HeadAccountID, Services: append([]string(nil), d.Services...), ActiveAccounts: counts[d.ID]})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Code < out[j].Code })
	return out
}
