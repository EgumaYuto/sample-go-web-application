package main

import (
	"net/http"

	"cabos.io/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()

	mux.GET("/health", controller.HealthController{}.GetHealthStatus)
	mux.PUT("/todo", controller.TodoController{}.AddTodo)
	mux.GET("/todo", controller.TodoController{}.GetTodoList)
	mux.GET("/todo/:id", controller.GetTodo)
	mux.DELETE("/todo/:id", controller.DeleteTodo)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
