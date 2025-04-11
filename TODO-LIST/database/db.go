package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	C "project/todo-list/config"
	"time"
)

func ConnectDatabase() (*sql.DB, error) {
	fmt.Println("Connect to database!")
	DB_username := os.Getenv("DB_USERNAME")
	DB_userpassword := os.Getenv("DB_USERPASSWORD")
	dsn := DB_username + ":" + DB_userpassword + "@tcp(127.0.0.1:3306)/TODO_LIST"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	err = db.Ping()
	fmt.Println("Ping vào database")
	if err != nil {
		fmt.Println("Ping thất bại")
		db.Close()
		return nil, err
	}
	C.Result("Ping")
	return db, nil
}
