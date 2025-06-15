package main

import (
	"log"
	"net/http"
	"time"
	"flag"
)
func main(){
	mux := http.NewServeMux()
	dirPath := flag.String("path", "./", "path of src to live-server")
	flag.Parse()
	fileServe := http.FileServer(http.Dir(*dirPath))
	mux.Handle("/",fileServe)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		WriteTimeout: 10 * time.Second ,
		ReadTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server running at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server failed: ", err)
	}
}

