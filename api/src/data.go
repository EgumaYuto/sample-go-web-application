package main

import (
	"database/sql"
	"log"

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

func insertTodo(title string) (todo Todo, err error) {
	insertStatement := "INSERT INTO TODO (TITLE) VALUES ( ? )"
	insertStmt, err := Db.Prepare(insertStatement)
	if err != nil {
		log.Println("prepase failure: ", err)
		return
	}
	defer insertStmt.Close()

	res, err := insertStmt.Exec(title)
	if err != nil {
		log.Println("exec failure: ", err)
		return
	}
	todo.Id, err = res.LastInsertId()
	return
}
