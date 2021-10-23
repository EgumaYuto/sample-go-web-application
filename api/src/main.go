package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func health(w http.ResponseWriter, r *http.Request) {
	_, err := Db.Exec("select 1")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error!!")
	} else {
		fmt.Fprintf(w, "OK!")
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	server.ListenAndServe()
}
