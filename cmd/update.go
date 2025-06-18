package cmd

import (
	"encoding/json"
	"fmt"
	"sample-todo-app/models"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [completada]",
	Short: "Actualizar el estado de una tarea",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: el ID debe ser un número entero\n")
			return
		}
		isComplete, err := strconv.ParseBool(args[1])
		if err != nil {
			fmt.Printf("Error: completada debe ser true/false\n")
			return
		}
		updatedTodo := map[string]interface{}{
			"is_complete": isComplete,
		}
		data, _, err := client.From("todos").Update(updatedTodo, "", "").Eq("id", strconv.Itoa(id)).Execute()
		if err != nil {
			fmt.Printf("Error al actualizar tarea: %v\n", err)
			return
		}
		var updated []models.Todo
		json.Unmarshal(data, &updated)
		if err != nil {
			fmt.Printf("Error al deserializar: %v\n", err)
			return
		}
		fmt.Printf("Tarea actualizada: %+v\n", updated[0])
	},
}
