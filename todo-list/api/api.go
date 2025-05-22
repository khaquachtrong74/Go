package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	C "project/todo-list/config"
	store "project/todo-list/database"
	"strconv"
)

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Get Tasks!")
	var tasks []string
	var err error
	tasks, err = store.GetData("SELECT task, state FROM tasks")
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
func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskIDInt, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		C.Failed("Convert int to string!")
		return }

	task, err := store.GetOneRowData("SELECT task, state FROM tasks WHERE id=?", taskIDInt)
	if err != nil {
		C.Failed("Get task!")
		return
	}
	if task == "" {
		C.Failed("Search task")
	}
	fmt.Fprintf(w, "<h1> Task Mà bạn kiếm nè</h1>"+"<h2>"+task+"</h2>")
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
	store.InsertData(&task, db)
	C.Result("Handle Post task!")
	HandleGetTasks(w, r)
	defer db.Close()
}
