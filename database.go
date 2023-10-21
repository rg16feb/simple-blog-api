package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password!@tcp(localhost:3306)/simple_blog")
	if err != nil {
		return nil, err
	}
	return db, nil
}
