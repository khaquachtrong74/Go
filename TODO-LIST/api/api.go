package api

import (
	"fmt"
	"html/template"
	"net/http"
	C "project/todo-list/config"
	store "project/todo-list/database"
)

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Get Tasks!")
	var tasks []string
	var err error
	tasks, err = store.GetData()
	if err != nil {
		fmt.Println("Get data failed!")
		return
	}
	tmpl, err := template.ParseFiles("./src/tasks.html")
	if err != nil {
		C.Failed("PArse Files")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, tasks)
	if err != nil {
		C.Failed("Template Execute")
	}

}
func HandlePostTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Post tasks")
	db, err := store.ConnectDatabase()
	if err != nil {
		C.Failed("Connect to database!")
		return
	}
	var task C.ToDo
	task.TASK = r.PostFormValue("task")
	task.STATE = r.PostFormValue("state")
	store.InsertData(&task, db)
	C.Result("Handle Post task!")
	HandleGetTasks(w, r)
	defer db.Close()
}
