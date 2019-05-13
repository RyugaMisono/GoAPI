package main

import (
	"fmt"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Todos list")
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Successfully created new Todo")
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Successfully updated Todo")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Destroyed that disgusting Todo")
}