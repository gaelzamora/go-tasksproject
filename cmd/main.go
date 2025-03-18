package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gaelzamora/go-rest-crud/config"
	"github.com/gaelzamora/go-rest-crud/internal/adapters/database"
	"github.com/gaelzamora/go-rest-crud/internal/adapters/handlers"
	"github.com/gaelzamora/go-rest-crud/internal/application"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.ConnectDB()

	userRepo := database.NewUserRepository(db)
	authService := application.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	taskRepo := database.NewTaskRepository(db)
	taskService := application.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := mux.NewRouter()
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(handlers.AuthMiddleware)
	protected.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	protected.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	protected.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	protected.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	protected.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")

	fmt.Println("🔵 Servidor corriendo en http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
