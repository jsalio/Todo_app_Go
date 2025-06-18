package cmd

import (
	"encoding/json"
	"fmt"
	"sample-todo-app/models"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [tarea]",
	Short: "Agregar una nueva tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		newTodo := models.Todo{
			Task:       task,
			IsComplete: false,
			CreateAt:   time.Now(),
		}
		todoMap := map[string]interface{}{
			"task":        newTodo.Task,
			"is_complete": newTodo.IsComplete,
			"created_at":  newTodo.CreateAt,
		}
		data, _, err := client.From("todos").Insert(todoMap, false, "*", "", "").Execute()
		if err != nil {
			fmt.Printf("Error al agregar tarea: %v\n", err)
			return
		}
		var insertedTodo []models.Todo
		json.Unmarshal(data, &insertedTodo)
		if err != nil {
			fmt.Printf("Error al deserializar: %v\n", err)
			return
		}
		fmt.Printf("Tarea agregada: %+v\n", insertedTodo[0])
	},
}
