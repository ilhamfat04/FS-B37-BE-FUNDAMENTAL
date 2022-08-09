package routes

import (
	// Import dumbmerch/handlers here ...
	h "dumbmerch/handlers"

	"github.com/gorilla/mux"
)

func TodoRoutes(r *mux.Router) {

	// Create Routes here ...
	r.HandleFunc("/todos", h.FindTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", h.GetTodo).Methods("GET")
	r.HandleFunc("/todo", h.CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", h.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todo/{id}", h.DeleteTodo).Methods("DELETE")
}
