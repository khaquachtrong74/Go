module task/database

go 1.24.3

replace task/database/operations => ./operations/

require task/database/operations v0.0.0-00010101000000-000000000000

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.33.0 // indirect
)
