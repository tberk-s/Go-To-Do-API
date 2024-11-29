package todo_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/tberk-s/Go-To-Do-API/internal/db"
	"github.com/tberk-s/Go-To-Do-API/internal/todo"
)

// MockDB ...
type MockDB struct {
	items []db.Item
}

// InsertItem ...
func (m *MockDB) InsertItem(ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

// GetAllItems ...
func (m *MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	return m.items, nil
}

// DeleteItem ...
func (m *MockDB) DeleteItem(ctx context.Context, task string) error {
	for i, item := range m.items {
		if item.Task == task {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return nil
		}
	}
	return nil
}

// Test for search.
func TestService_Search(t *testing.T) {
	tests := []struct {
		name       string
		todosToAdd []string
		query      string
		expected   []string
	}{
		{
			name:       "todo is shop, query is `sh` we expect shop",
			todosToAdd: []string{"shop"},
			query:      "sh",
			expected:   []string{"shop"},
		},
		{
			name:       "todo is Shopping, query is `sh` we expect Shopping",
			todosToAdd: []string{"Shopping"},
			query:      "sh",
			expected:   []string{"Shopping"},
		},
		{
			name:       "todo is SHOP, query is `sh` we expect SHOP",
			todosToAdd: []string{"SHOP"},
			query:      "sh",
			expected:   []string{"SHOP"},
		},
		{
			name:       "todo is go shop, query is `go ` we expect go shop",
			todosToAdd: []string{"go shop"},
			query:      "go ",
			expected:   []string{"go shop"},
		},
		{
			name:       "todo is go hophop, query is `go h` we expect go hophop",
			todosToAdd: []string{"go hophop"},
			query:      "go h",
			expected:   []string{"go hophop"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDB{}
			svc := todo.New(mockDB)

			for _, toAdd := range tt.todosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Fatalf("Add() error = %v", err)
				}
			}

			got, err := svc.Search(tt.query)
			if err != nil {
				t.Fatalf("Search() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Search() = %v, want %v", got, tt.expected)
			}
		})
	}
}
