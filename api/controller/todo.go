package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"cabos.io/model"
	"github.com/julienschmidt/httprouter"
)

func AddTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var todo model.Todo
	json.Unmarshal(body, &todo)

	var err error
	todo, err = model.InsertTodo(todo.Title)
	if err != nil {
		log.Println("add failure: ", err)
		w.WriteHeader(500)
		return
	}

	output, err := json.MarshalIndent(&todo, "", "\t\t")
	if err != nil {
		log.Println("marshal failure: ", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(output)
}

func GetTodoList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todos, err := model.GetTodoList()
	if err != nil {
		log.Println("list failure: ", err)
		w.WriteHeader(500)
		return
	}

	output, err := json.MarshalIndent(&todos, "", "\t\t")
	if err != nil {
		log.Println("marshal failure: ", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(output)
}

func GetTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Println("parse failure: ", err)
		w.WriteHeader(500)
		return
	}

	todo, err := model.GetTodoById(id)

	if err != nil {
		log.Println("get failure: ", err)
		w.WriteHeader(500)
		return
	}

	output, err := json.MarshalIndent(&todo, "", "\t\t")
	if err != nil {
		log.Println("marshal failure: ", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(output)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Println("parse failure: ", err)
		w.WriteHeader(500)
		return
	}

	model.DeleteTodoById(id)

	todo := model.Todo{Id: int64(id)}
	output, err := json.MarshalIndent(&todo, "", "\t\t")
	if err != nil {
		log.Println("marshal failure: ", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(output)
}
