package ports

import "sample-todo-app/models"

// TodoUseCases define las operaciones de negocio para las tareas
type TodoUseCases interface {
	CreateTask(task string) (*models.Todo, error)
	GetTask(id int) (*models.Todo, error)
	ListTasks() ([]models.Todo, error)
	UpdateTask(id int, task string, isComplete bool) (*models.Todo, error)
	DeleteTask(id int) error
}
