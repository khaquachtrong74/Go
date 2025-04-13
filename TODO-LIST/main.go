package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	api "project/todo-list/api"
)

//func OpenServer(mux *http.ServeMux) {
//	fmt.Println("Open Server!")
//	s := http.Server{
//		Addr:         ":8080",
//		Handler:      mux,
//		ReadTimeout:  15 * time.Second,
//		WriteTimeout: 15 * time.Second,
//	}
//	log.Fatal(s.ListenAndServe())
//}

func main() {
	fmt.Println("Start backend")
	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/{id}",api.HandleGetTask)
	//	mux.HandleFunc("/", api.HandleGetTasks)
	//	mux.HandleFunc("/add-task", api.HandlePostTasks)
	r := chi.NewRouter()
	r.Get("/tasks/{id}", api.HandleGetTask)
	r.Get("/tasks", api.HandleGetTasks)
	r.Post("/add-task", api.HandlePostTasks)
	log.Fatal(http.ListenAndServe(":8080", r))
	//	OpenServer(mux)

}
