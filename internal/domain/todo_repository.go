package domain

type TodoRepository interface {
	Create(todo *Todo) error
	GetAll() ([]*Todo, error)
}
