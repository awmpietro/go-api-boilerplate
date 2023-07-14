package main

import (
	"log"
	"net/http"
	"os"

	"teste-go/internal/db"
	"teste-go/internal/delivery"
	"teste-go/internal/repository"
	"teste-go/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize database connection
	db.InitPostgreSQL()
	defer db.ClosePostgreSQL()

	// Create repositories
	userRepository := repository.NewUserRepository(db.GetDB())
	productRepository := repository.NewProductRepository(db.GetDB())

	// Create use cases
	userUseCase := usecase.NewUserUseCase(userRepository)
	productUseCase := usecase.NewProductUseCase(productRepository)

	// Create handlers
	userHandler := delivery.NewUserHandler(userUseCase)
	productHandler := delivery.NewProductHandler(productUseCase)

	// Create multiplexer (router)
	mux := http.NewServeMux()

	// Register routes
	delivery.RegisterUserRoutes(userHandler, mux)
	delivery.RegisterProductRoutes(productHandler, mux)

	// Start the server
	log.Printf("Listening on port %s", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), mux))
}
