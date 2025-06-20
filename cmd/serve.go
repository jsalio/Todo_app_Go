package cmd

import (
	"fmt"
	"log"
	"os"
	"sample-todo-app/internal/application/usecases"
	"sample-todo-app/internal/infrastructure/repository"
	"sample-todo-app/internal/interfaces/api"

	"github.com/spf13/cobra"
	"github.com/supabase-community/supabase-go"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia el servidor web de la aplicación",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Asegurarse de que las variables de entorno estén cargadas
		if !PassPrerequirements() {
			os.Exit(1)
		}

		// Inicializar cliente de Supabase
		client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
		if err != nil {
			log.Fatalf("Error al inicializar cliente de Supabase: %v", err)
		}

		// Inicializar el repositorio y los casos de uso
		repo := repository.NewSupabaseTodoRepository(client)
		todoUseCases = usecases.NewTodoUseCases(repo)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Crear e iniciar el servidor web
		server := api.NewServer(todoUseCases)
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		fmt.Printf("Servidor iniciado en http://localhost:%s\n", port)
		fmt.Println("Presiona Ctrl+C para detener el servidor")

		if err := server.Start(":" + port); err != nil {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
