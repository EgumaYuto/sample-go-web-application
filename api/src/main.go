package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func health(w http.ResponseWriter, r *http.Request) {
	_, err := healthCheck()
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
