package main

import (
	"net/http"

	"cabos.io/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()

	mux.GET("/health", controller.HealthController{}.GetHealthStatus)
	mux.PUT("/todo", controller.AddTodo)
	mux.GET("/todo", controller.GetTodoList)
	mux.GET("/todo/:id", controller.GetTodo)
	mux.DELETE("/todo/:id", controller.DeleteTodo)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
