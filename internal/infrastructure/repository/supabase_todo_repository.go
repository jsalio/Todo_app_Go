package repository

import (
	"strconv"

	"sample-todo-app/internal/domain/ports"
	"sample-todo-app/models"

	"github.com/supabase-community/supabase-go"
)

type supabaseTodoRepository struct {
	client *supabase.Client
}

// NewSupabaseTodoRepository crea un nuevo repositorio de tareas usando Supabase
func NewSupabaseTodoRepository(client *supabase.Client) ports.TodoRepository {
	return &supabaseTodoRepository{client: client}
}

func (r *supabaseTodoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	var result models.Todo

	// Creamos un mapa con solo los campos que queremos insertar
	// para evitar enviar el ID que debe ser generado por la base de datos
	insertData := map[string]interface{}{
		"task":        todo.Task,
		"is_complete": todo.IsComplete,
		"created_at":  todo.CreatedAt,
	}

	// Usamos Insert con el parámetro returning=representation para obtener el registro insertado
	_, err := r.client.From("todos").
		Insert(insertData, false, "", "representation", "").
		Single().
		ExecuteTo(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *supabaseTodoRepository) GetByID(id int) (*models.Todo, error) {
	var todo models.Todo
	// Usamos Select con eq para filtrar por ID
	_, err := r.client.From("todos").Select("*", "exact", false).Eq("id", strconv.Itoa(id)).Single().ExecuteTo(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *supabaseTodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	// Obtenemos todos los registros
	_, err := r.client.From("todos").Select("*", "exact", false).ExecuteTo(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *supabaseTodoRepository) Update(todo *models.Todo) (*models.Todo, error) {
	var result models.Todo
	// Usamos Update con el parámetro returning=representation para obtener el registro actualizado
	_, err := r.client.From("todos").Update(todo, "", "").Eq("id", strconv.Itoa(todo.ID)).
		ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *supabaseTodoRepository) Delete(id int) error {
	// Usamos Delete para eliminar el registro
	_, _, err := r.client.From("todos").Delete("", "").Eq("id", strconv.Itoa(id)).Execute()
	return err
}
