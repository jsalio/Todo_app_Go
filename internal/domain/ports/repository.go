package ports

import "sample-todo-app/models"

// TodoRepository define la interfaz para el repositorio de tareas
type TodoRepository interface {
	Create(todo *models.Todo) (*models.Todo, error)
	GetByID(id int) (*models.Todo, error)
	GetAll() ([]models.Todo, error)
	Update(todo *models.Todo) (*models.Todo, error)
	Delete(id int) error
}
