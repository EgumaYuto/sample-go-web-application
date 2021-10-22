package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This service is healthy.")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	server.ListenAndServe()
}
