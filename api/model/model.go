package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	user := "root"
	password := "example"
	protocol := "tcp(db:3306)"
	dbName := "test_db"
	db, err = sql.Open("mysql", user+":"+password+"@"+protocol+"/"+dbName)
	if err != nil {
		panic(err)
	}
}

func HealthCheck() (sql.Result, error) {
	return db.Exec("select 1")
}
