module project/todo-list/api

go 1.24.1

replace project/todo-list/database => ../database

replace project/todo-list/config => ../config

require (
	github.com/go-chi/chi/v5 v5.2.1
	project/todo-list/config v0.0.0-00010101000000-000000000000
	project/todo-list/database v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
)
