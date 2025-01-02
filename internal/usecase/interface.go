package usecase

import "todo-api/internal/domain/models"

type TodoUseCase interface {
	GetTodos() ([]models.Todo, error)
	GetTodo(id string) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	UpdateTodo(id string, todo models.Todo) (models.Todo, error)
	DeleteTodo(id string) error
}
