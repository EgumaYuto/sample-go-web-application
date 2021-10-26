package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func addTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var todo Todo
	json.Unmarshal(body, &todo)

	insertStatement := "INSERT INTO TODO (TITLE) VALUES ( ? )"
	insertStmt, err := Db.Prepare(insertStatement)
	if err != nil {
		w.WriteHeader(500)
		log.Println("prepase failure: ", err)
		return
	}
	defer insertStmt.Close()

	res, err := insertStmt.Exec(todo.Title)
	if err != nil {
		w.WriteHeader(500)
		log.Println("exec failure: ", err)
		return
	}
	todo.Id, _ = res.LastInsertId()

	w.WriteHeader(200)
}
