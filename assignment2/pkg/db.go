package pkg

import (
	"database/sql"	
	_ "github.com/lib/pq"
	"github.com/AminaAmangeldi15/Go/assignment2"
)

func Db() *sql.DB{
	connStr := "user=postgres password=9792amina dbname=golang sslmode=disable"
    DB, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer DB.Close()
	return DB
}