package persistence

import (
	"errors"

	"github.com/PlasmaXD/CleanArchitecture/internal/domain"
)

type todoRepository struct {
	store  []*domain.Todo
	nextID int64
}

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{store: make([]*domain.Todo, 0), nextID: 1}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	if todo.Title == "" {
		return errors.New("title is empty")
	}
	todo.ID = r.nextID
	r.nextID++
	r.store = append(r.store, todo)
	return nil
}

func (r *todoRepository) GetAll() ([]*domain.Todo, error) {
	return r.store, nil
}
