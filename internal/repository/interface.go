package repository

import "todo-api/internal/domain/models"

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetByID(id string) (models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(id string, todo models.Todo) (models.Todo, error)
	Delete(id string) error
}
