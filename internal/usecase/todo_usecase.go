package usecase

import (
	"time"
	"todo-api/internal/domain/models"
	"todo-api/internal/repository"
)

type todoUseCase struct {
	repo repository.TodoRepository
}

func NewTodoUseCase(repo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		repo: repo,
	}
}

func (u *todoUseCase) GetTodos() ([]models.Todo, error) {
	return u.repo.GetAll()
}

func (u *todoUseCase) GetTodo(id string) (models.Todo, error) {
	return u.repo.GetByID(id)
}

func (u *todoUseCase) CreateTodo(todo models.Todo) (models.Todo, error) {
	todo.CreatedAt = time.Now()
	return u.repo.Create(todo)
}

func (u *todoUseCase) UpdateTodo(id string, todo models.Todo) (models.Todo, error) {
	return u.repo.Update(id, todo)
}

func (u *todoUseCase) DeleteTodo(id string) error {
	return u.repo.Delete(id)
}
