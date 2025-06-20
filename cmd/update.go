package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [tarea] [completada]",
	Short: "Actualizar una tarea",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: el ID debe ser un número entero\n")
			return
		}

		task := args[1]

		isComplete, err := strconv.ParseBool(args[2])
		if err != nil {
			fmt.Printf("Error: completada debe ser true/false\n")
			return
		}

		todo, err := todoUseCases.UpdateTask(id, task, isComplete)
		if err != nil {
			fmt.Printf("Error al actualizar tarea: %v\n", err)
			return
		}

		fmt.Printf("Tarea actualizada: ID: %d, Tarea: %s, Completada: %v\n",
			todo.ID, todo.Task, todo.IsComplete)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
