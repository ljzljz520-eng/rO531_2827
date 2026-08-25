package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"hospitalportal/internal/domain"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var ErrClosed = errors.New("store closed")
var accountsBucket = []byte("accounts")
var departmentsBucket = []byte("departments")
var shiftsBucket = []byte("shifts")
var auditsBucket = []byte("audits")
var metadataBucket = []byte("metadata")

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	if path == "" {
		return nil, fmt.Errorf("path required")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, fmt.Errorf("create store directory: %w", err)
	}
	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	s := &Store{db: db, path: path}
	if err := s.migrate(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) migrate() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, name := range [][]byte{accountsBucket, departmentsBucket, shiftsBucket, auditsBucket, metadataBucket} {
			if _, err := tx.CreateBucketIfNotExists(name); err != nil {
				return fmt.Errorf("create bucket %s: %w", name, err)
			}
		}
		m := tx.Bucket(metadataBucket)
		if m.Get([]byte("schema_version")) == nil {
			if err := m.Put([]byte("schema_version"), []byte("1")); err != nil {
				return err
			}
		}
		return nil
	})
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func (s *Store) View(fn func(*bbolt.Tx) error) error {
	if s == nil || s.db == nil {
		return ErrClosed
	}
	return s.db.View(fn)
}
func (s *Store) Update(fn func(*bbolt.Tx) error) error {
	if s == nil || s.db == nil {
		return ErrClosed
	}
	return s.db.Update(fn)
}
func putJSON(tx *bbolt.Tx, bucket []byte, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("encode %s: %w", key, err)
	}
	if err := tx.Bucket(bucket).Put([]byte(key), data); err != nil {
		return fmt.Errorf("write %s: %w", key, err)
	}
	return nil
}
func getJSON(tx *bbolt.Tx, bucket []byte, key string, out any, notFound error) error {
	data := tx.Bucket(bucket).Get([]byte(key))
	if data == nil {
		return notFound
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("decode %s: %w", key, err)
	}
	return nil
}
func listJSON[T any](tx *bbolt.Tx, bucket []byte, match func(T) bool) ([]T, error) {
	values := []T{}
	err := tx.Bucket(bucket).ForEach(func(_, data []byte) error {
		var value T
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		if match == nil || match(value) {
			values = append(values, value)
		}
		return nil
	})
	return values, err
}
func sortAccounts(v []domain.UserAccount) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].DisplayName == v[j].DisplayName {
			return v[i].ID < v[j].ID
		}
		return v[i].DisplayName < v[j].DisplayName
	})
}
func sortDepartments(v []domain.Department) {
	sort.Slice(v, func(i, j int) bool { return v[i].Code < v[j].Code })
}
func sortShifts(v []domain.DutyShift) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].StartAt.Equal(v[j].StartAt) {
			return v[i].ID < v[j].ID
		}
		return v[i].StartAt.Before(v[j].StartAt)
	})
}
func sortAudits(v []domain.AuditRecord) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].OccurredAt.Equal(v[j].OccurredAt) {
			return v[i].ID < v[j].ID
		}
		return v[i].OccurredAt.Before(v[j].OccurredAt)
	})
}
