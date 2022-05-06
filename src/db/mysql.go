package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql",
		"root:tanm146@tcp(127.0.0.1:3306)/douyin")
	if err != nil {
		panic(err)
	}
	return db
}
