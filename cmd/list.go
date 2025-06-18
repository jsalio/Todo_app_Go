package cmd

import (
	"encoding/json"
	"fmt"
	"sample-todo-app/models"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas las tareas",
	Run: func(cmd *cobra.Command, args []string) {
		data, _, err := client.From("todos").Select("*", "exact", false).Execute()
		if err != nil {
			fmt.Printf("Error al listar tareas: %v\n", err)
			return
		}
		var todos []models.Todo
		json.Unmarshal(data, &todos)
		if err != nil {
			fmt.Printf("Error al deserializar: %v\n", err)
			return
		}
		if len(todos) == 0 {
			fmt.Println("No hay tareas")
			return
		}
		for _, todo := range todos {
			fmt.Printf("ID: %d, Tarea: %s, Completada: %v, Creada: %v\n",
				todo.ID, todo.Task, todo.IsComplete, todo.CreateAt)
		}
	},
}
