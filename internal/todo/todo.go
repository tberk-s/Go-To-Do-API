package todo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tberk-s/Go-To-Do-API/internal/db"
)

// Manager ...
type Manager interface {
	InsertItem(ctx context.Context, item db.Item) error
	GetAllItems(ctx context.Context) ([]db.Item, error)
	DeleteItem(ctx context.Context, task string) error
}

// Service ...
type Service struct {
	db Manager
}

// Item ...
type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

// New ...
func New(database Manager) *Service {
	return &Service{
		db: database,
	}
}

// Add ...
func (svc *Service) Add(todo string) error {
	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("failed to read from db: %w", err)
	}
	for _, t := range items {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}

	return nil
}

// GetAll ...
func (svc *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read from db: %w", err)
	}

	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}

	return results, nil
}

// Search ...
func (svc *Service) Search(query string) ([]string, error) {
	var result []string
	items, err := svc.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read from db: %w", err)
	}

	for _, todo := range items {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			result = append(result, todo.Task)
		}
	}

	return result, nil
}

// Delete ...
func (svc *Service) Delete(task string) error {
	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("failed to read from db: %w", err)
	}

	taskExists := false
	for _, item := range items {
		if item.Task == task {
			taskExists = true

			break
		}
	}

	if !taskExists {
		return fmt.Errorf("task '%s' does not exist", task)
	}

	if err := svc.db.DeleteItem(context.Background(), task); err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return svc.db.DeleteItem(context.Background(), task)
}
