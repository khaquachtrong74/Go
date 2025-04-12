package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	api "project/todo-list/api"
	"time"
)

func OpenServer(mux *http.ServeMux) {
	fmt.Println("Open Server!")
	s := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
func main() {
	fmt.Println("Start backend")
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HandleGetTasks)
	mux.HandleFunc("/add-task", api.HandlePostTasks)
	OpenServer(mux)
}
