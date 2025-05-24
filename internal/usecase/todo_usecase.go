package usecase

import (
	"github.com/PlasmaXD/CleanArchitecture/internal/domain"
)

type TodoUseCase interface {
	CreateTodo(title string) (*domain.Todo, error)
	ListTodos() ([]*domain.Todo, error)
}

type todoUseCase struct {
	repo domain.TodoRepository
}

func NewTodoUseCase(r domain.TodoRepository) TodoUseCase {
	return &todoUseCase{repo: r}
}

func (u *todoUseCase) CreateTodo(title string) (*domain.Todo, error) {
	todo := &domain.Todo{Title: title}
	if err := u.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUseCase) ListTodos() ([]*domain.Todo, error) {
	return u.repo.GetAll()
}
