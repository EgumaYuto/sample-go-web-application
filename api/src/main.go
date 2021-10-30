package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello!")
}

func health(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := healthCheck()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error!!")
	} else {
		fmt.Fprintf(w, "OK!")
	}
}

func main() {
	mux := httprouter.New()

	mux.GET("/health", health)
	mux.GET("/hello", hello)
	mux.PUT("/todo", addTodo)
	mux.GET("/todo", getTodoList)
	mux.GET("/todo/:id", getTodo)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
