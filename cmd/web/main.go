package main

import (
	"fmt"
	"log"
	"os"

	"sample-todo-app/internal/application/usecases"
	"sample-todo-app/internal/infrastructure/repository"
	"sample-todo-app/internal/interfaces/api"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando archivo .env: %v", err)
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL y SUPABASE_KEY deben estar configurados en el archivo .env")
	}

	// Inicializar cliente de Supabase
	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		log.Fatalf("Error al inicializar cliente de Supabase: %v", err)
	}

	// Inicializar el repositorio y los casos de uso
	repo := repository.NewSupabaseTodoRepository(client)
	todoUseCase := usecases.NewTodoUseCases(repo)

	// Crear e iniciar el servidor web
	server := api.NewServer(todoUseCase)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor iniciado en http://localhost:%s\n", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
