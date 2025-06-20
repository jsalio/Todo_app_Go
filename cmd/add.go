package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [tarea]",
	Short: "Agregar una nueva tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		todo, err := todoUseCases.CreateTask(task)
		if err != nil {
			fmt.Printf("Error al agregar tarea: %v\n", err)
			return
		}
		fmt.Printf("Tarea agregada: ID: %d, Tarea: %s, Completada: %v\n", 
			todo.ID, todo.Task, todo.IsComplete)
	},
}
