package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)
func main(){
	// Open my.db data file in current dir
	// Created it if it doesn't exit
	
	// OK but it will wait until the other
	// process closes.
	//db, err := bolt.Open("my.db", 0600, nil)
	
	// Try this
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
}
