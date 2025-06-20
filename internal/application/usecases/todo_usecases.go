package usecases

import (
	"time"

	"sample-todo-app/internal/domain/ports"
	"sample-todo-app/models"
)

type todoUseCases struct {
	repo ports.TodoRepository
}

// NewTodoUseCases crea una nueva instancia de TodoUseCases
func NewTodoUseCases(repo ports.TodoRepository) ports.TodoUseCases {
	return &todoUseCases{repo: repo}
}

func (uc *todoUseCases) CreateTask(task string) (*models.Todo, error) {
	todo := &models.Todo{
		Task:       task,
		IsComplete: false,
		CreatedAt:  time.Now(),
	}
	return uc.repo.Create(todo)
}

func (uc *todoUseCases) GetTask(id int) (*models.Todo, error) {
	return uc.repo.GetByID(id)
}

func (uc *todoUseCases) ListTasks() ([]models.Todo, error) {
	return uc.repo.GetAll()
}

func (uc *todoUseCases) UpdateTask(id int, task string, isComplete bool) (*models.Todo, error) {
	todo, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	todo.Task = task
	todo.IsComplete = isComplete
	return uc.repo.Update(todo)
}

func (uc *todoUseCases) DeleteTask(id int) error {
	return uc.repo.Delete(id)
}
