module project/login

go 1.24.1

replace project/login/database => ./database

require project/login/database v0.0.0-00010101000000-000000000000

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
)
