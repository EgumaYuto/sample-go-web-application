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

	var err error
	todo, err = insertTodo(todo.Title)
	if err != nil {
		log.Println("add failure: ", err)
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}
