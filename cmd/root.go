package cmd

import (
	"fmt"
	"os"

	"sample-todo-app/internal/application/usecases"
	"sample-todo-app/internal/domain/ports"
	supabaseRepo "sample-todo-app/internal/infrastructure/repository"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/supabase-community/supabase-go"
)

var (
	supabaseURL  string
	supabaseKey  string
	todoUseCases ports.TodoUseCases
)

// rootCmd representa el comando base
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "CLI para gestionar tareas con Supabase",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		errEnv := godotenv.Load()

		if errEnv != nil {
			fmt.Printf("Error cargando archivo .env: %v\n", errEnv)
			os.Exit(1)
		}

		supabaseURL = os.Getenv("SUPABASE_URL")
		supabaseKey = os.Getenv("SUPABASE_KEY")

		if supabaseURL == "" || supabaseKey == "" {
			fmt.Println("Error: SUPABASE_URL o SUPABASE_KEY no están definidas en el archivo .env")
			os.Exit(1)
		}

		// Inicializar cliente de Supabase
		client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
		if err != nil {
			fmt.Printf("Error al inicializar cliente de Supabase: %v\n", err)
			os.Exit(1)
		}

		// Inicializar el repositorio y los casos de uso
		repo := supabaseRepo.NewSupabaseTodoRepository(client)
		todoUseCases = usecases.NewTodoUseCases(repo)
	},
}

// Execute ejecuta el comando raíz
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
