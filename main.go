package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! I'm ok")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello).Methods("GET")
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", UpdateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))
}