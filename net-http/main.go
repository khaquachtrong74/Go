package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/dreamsofcode-io/nethttp/monster"
	"github.com/khaquachtrong74/simple-api/net-http/middleware"
)

func main(){
	handler := &monster.Handler{}
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", func(w http.ResponseWriter, r *http.Request){
		// something done here
		// path parameters
		id := r.PathValue("id")
		w.Write([] byte("Received request for the item: " + id + "\n"))
	})
	// method based routing
	router.HandleFunc("GET /monster", handler.Create)
	router.HandleFunc("PUT /monster", handler.UpdateByID)
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Starting server on port: " + s.Addr)
	log.Fatal(s.ListenAndServe())
}
