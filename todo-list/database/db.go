package database

import (
	"database/sql"
	"fmt"
	"os"
	C "project/todo-list/config"
	"time"
	_ "github.com/go-sql-driver/mysql"
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
func GetData(query string) ([]string, error) {
	db, err := ConnectDatabase()
	if err != nil {
		C.Failed("GET TASKS")
		return nil, err
	}
	rows, err := db.Query(query)
	if err != nil {
		C.Failed("Querying")
		return nil, err
	}
	var tasks []string
	for rows.Next() {
		var task C.ToDo
		if err := rows.Scan(&task.TASK, &task.STATE); err != nil {
			C.Failed("Read list")
			return nil, err
		}
		var str string = task.TASK + ":" + task.STATE
		tasks = append(tasks, str)
	}
	return tasks, err
}
func GetOneRowData(query string, id int) (string, error) {
	db, err := ConnectDatabase()
	if err != nil {
		C.Failed("GET TASKS")
		return "", err
	}
	var task string
	var state string
	fmt.Println(query)
	fmt.Println(id)
	err = db.QueryRow(query, id).Scan(&task, &state)
	if err != nil {
		C.Failed("Querying")
		return "", err
	}
	return task + " : " + state, err
}
func InsertData(data *C.ToDo, db *sql.DB) {
	fmt.Println("Xử lý Update data/Chèn")
	if data.TASK == "" {
		fmt.Println("Không có task nên không lưu")
		return
	}
	var insertQuery string
	insertQuery = "INSERT INTO tasks(task) VALUES(?)"
	db.Exec(insertQuery, data.TASK)
	fmt.Println("Hoàn thành xử lý chèn ")
}
