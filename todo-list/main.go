package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	api "project/todo-list/api"
)
func main() {
	fmt.Println("Start backend")
	r := chi.NewRouter()
	r.Get("/tasks/{id}", api.HandleGetTask)
	r.Get("/tasks", api.HandleGetTasks)
	r.Post("/add-task", api.HandlePostTasks)
	log.Fatal(http.ListenAndServe(":8080", r))
}
