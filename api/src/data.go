package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	user := "root"
	password := "example"
	protocol := "tcp(db:3306)"
	dbName := "test_db"
	Db, err = sql.Open("mysql", user+":"+password+"@"+protocol+"/"+dbName)
	if err != nil {
		panic(err)
	}
}

func healthCheck() (sql.Result, error) {
	return Db.Exec("select 1")
}
