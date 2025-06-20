package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Eliminar una tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: el ID debe ser un número entero\n")
			return
		}

		err = todoUseCases.DeleteTask(id)
		if err != nil {
			fmt.Printf("Error al eliminar tarea: %v\n", err)
			return
		}
		fmt.Println("Tarea eliminada con ID:", id)
	},
}
