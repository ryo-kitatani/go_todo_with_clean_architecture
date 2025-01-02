package mysql

import (
	"todo-api/internal/domain/models"
	"todo-api/internal/repository"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repository.TodoRepository {
	db.AutoMigrate(&models.Todo{})
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *mysqlRepository) GetByID(id string) (models.Todo, error) {
	var todo models.Todo
	result := r.db.First(&todo, id)
	return todo, result.Error
}

func (r *mysqlRepository) Create(todo models.Todo) (models.Todo, error) {
	result := r.db.Create(&todo)
	return todo, result.Error
}

func (r *mysqlRepository) Update(id string, todo models.Todo) (models.Todo, error) {
	var existingTodo models.Todo
	if err := r.db.First(&existingTodo, id).Error; err != nil {
		return models.Todo{}, err
	}

	result := r.db.Model(&existingTodo).Updates(todo)
	return existingTodo, result.Error
}

func (r *mysqlRepository) Delete(id string) error {
	result := r.db.Delete(&models.Todo{}, id)
	return result.Error
}
