package api

import (
	"fmt"

	"sample-todo-app/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	todoUseCase ports.TodoUseCases
}

func NewServer(todoUseCase ports.TodoUseCases) *Server {
	server := &Server{
		todoUseCase: todoUseCase,
	}
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	s.router = gin.Default()

	// Configurar rutas de la API
	api := s.router.Group("/api")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", s.listTodos)
			todos.POST("", s.createTodo)
			todos.GET("/:id", s.getTodo)
			todos.PUT("/:id", s.updateTodo)
			todos.DELETE("/:id", s.deleteTodo)
		}
	}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// Handler para listar todos los TODOs
func (s *Server) listTodos(c *gin.Context) {
	todos, err := s.todoUseCase.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todos)
}

// Handler para crear un nuevo TODO
func (s *Server) createTodo(c *gin.Context) {
	var request struct {
		Task string `json:"task" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Solicitud inválida"})
		return
	}

	todo, err := s.todoUseCase.CreateTask(request.Task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, todo)
}

// Handler para obtener un TODO por ID
func (s *Server) getTodo(c *gin.Context) {
	id := c.Param("id")
	// Convertir id a int (necesitarás manejar el error aquí)
	// Por simplicidad, asumimos que el parámetro es un número válido
	// En una aplicación real, deberías validar esto
	var todoID int
	_, err := fmt.Sscanf(id, "%d", &todoID)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	todo, err := s.todoUseCase.GetTask(todoID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Tarea no encontrada"})
		return
	}

	c.JSON(200, todo)
}

// Handler para actualizar un TODO
func (s *Server) updateTodo(c *gin.Context) {
	id := c.Param("id")
	// Convertir id a int
	var todoID int
	_, err := fmt.Sscanf(id, "%d", &todoID)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var request struct {
		Task       *string `json:"task,omitempty"`
		IsComplete *bool   `json:"is_complete,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Solicitud inválida"})
		return
	}

	// Obtener la tarea actual para preservar los valores no proporcionados
	currentTodo, err := s.todoUseCase.GetTask(todoID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Tarea no encontrada"})
		return
	}

	// Actualizar solo los campos proporcionados
	task := currentTodo.Task
	if request.Task != nil {
		task = *request.Task
	}

	isComplete := currentTodo.IsComplete
	if request.IsComplete != nil {
		isComplete = *request.IsComplete
	}

	updatedTodo, err := s.todoUseCase.UpdateTask(todoID, task, isComplete)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedTodo)
}

// Handler para eliminar un TODO
func (s *Server) deleteTodo(c *gin.Context) {
	id := c.Param("id")
	// Convertir id a int
	var todoID int
	_, err := fmt.Sscanf(id, "%d", &todoID)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	err = s.todoUseCase.DeleteTask(todoID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
