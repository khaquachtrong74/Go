package main

import (
	"log"
	DB "todo-app/database/operations"
	"flag"
)
var(
	task = flag.String("task", "Default task", "Enter your task today!")
	priority = flag.Int("priority", 0, "Enter priority of task # default 0")
	read = flag.Bool("read",false,"Switch true to read task today bro! # `bool`")
	clean = flag.Bool("clear", false, "Switch true to clear boltDB # `bool`")
)
func main(){
	flag.Parse()
	db, err := DB.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Write(db,*task,*priority)
	if err != nil {
		log.Fatal("failed to write task")
	}
	/*
		check permissions of operation read and clean,
		Use func IsWork to track error when it happen on what operation
	*/
	if *read == true && DB.IsWork(DB.ReadDB(db), "Read data"){
		log.Print("read work")
	}
	// tasks is my dedfault bucket so if you wana change
	// please Override OpenDB() and change default bucketName
	if *clean == true &&	DB.IsWork(DB.DeleteBucket(db, []byte("tasks")), "Clear data")	{
		log.Print("clean work")
	}

	defer db.Close()
}
