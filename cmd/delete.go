package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Eliminar una tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		_, _, err := client.From("todos").Delete("", "").Eq("id", id).Execute()
		if err != nil {
			fmt.Printf("Error al eliminar tarea: %v\n", err)
			return
		}
		fmt.Println("Tarea eliminada con ID:", id)
	},
}
