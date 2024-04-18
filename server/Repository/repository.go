package repository

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository() *URLRepository {
	db, err := sql.Open("postgres", "user=myname dbname=mydb")
	if err != nil {
		panic(err)
	}
	return &URLRepository{DB: db}
}
