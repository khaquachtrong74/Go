package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	C "project/todo-list/config"
	store "project/todo-list/database"
	"time"
)

func HandleMuxServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Mux Server")
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Greeting, welcome to my place~~~and dieeee!"))
	case http.MethodPost:
	}
}
func handleGetTasks(w http.ResponseWriter, r *http.Request) {
	db, err := store.ConnectDatabase()
	if err != nil {
		C.Failed("GET TASKS")
		return
	}
	query := "SELECT id, task, state FROM tasks"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("ERROR: Querying tasks failed: %v\n", err)
		C.Failed("Querying")
		return
	}
	defer rows.Close()
	for rows.Next() {
		var task C.ToDo
		if err := rows.Scan(&task.ID, &task.TASK, &task.STATE); err != nil {
			C.Failed("Read list")
			return
		}
		fmt.Fprintf(w, string(task.ID)+":"+task.TASK+":"+task.STATE)
	}

	if err = rows.Err(); err != nil {
		log.Printf("ERROR: Error after iterating rows: %v\n", err)
		return
	}
	C.Result("Querying")
}

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
	db, err := store.ConnectDatabase()
	if err == nil {
		fmt.Println("Connect database success!")
	}
	_, err = db.Query("SELECT * FROM tasks")
	if err == nil {
		C.Result("Query")
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleMuxServer(w, r)
	})
	mux.HandleFunc("/tasks", handleGetTasks)
	OpenServer(mux)
}
