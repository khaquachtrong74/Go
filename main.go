package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"project/login/database"
)

func handlerWeb(w http.ResponseWriter, r *http.Request, path string) {
	http.ServeFile(w, r, path)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		http.ServeFile(w, r, "./public/index.html")
	case r.URL.Path == "/login":
		http.FileServer(http.Dir("./public/login"))
	default:
		http.NotFound(w, r) // return 404
	}
}

var db *sql.DB

func main() {
	var err error
	fmt.Println("Listen And Serve with Go!")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		database.ServeFormHandle(w, r, "./public/index.html")
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		database.SubmitLoginHandler(w, r, "./public/login.html")
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		database.SubmitRegisterHandle(w, r, "./public/register.html")
	})
	staticFolderPath := http.FileServer(http.Dir("./public/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFolderPath))
	sever := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	fmt.Println("Server starting on port :8080...")
	err = sever.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
