package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Declare Todos Struct here ...
type Todos struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"condition"`
}

// Declare Todos Global Variable ...
var todos = []Todos{
	{
		Id:     "1",
		Title:  "Cuci tangan",
		IsDone: true,
	},
	{
		Id:     "2",
		Title:  "Jaga jarak",
		IsDone: false,
	},
}

func main() {
	r := mux.NewRouter()

	// Create routes here ...
	r.HandleFunc("/todos", FindTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	r.HandleFunc("/todo", CreateTodo).Methods("POST")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

// Create FindTodos Function here ...
func FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// Create GetTodo Function here ...
func GetTodo(w http.ResponseWriter, r *http.Request) {

	// data1 := "ilham"
	// data2 := &data1

	// fmt.Println(data1)
	// fmt.Println(data2)

	w.Header().Set("Content-Type", "application/json")

	// get value form params
	params := mux.Vars(r)
	id := params["id"]

	// init variable to store data todo
	var todoData Todos
	var isGetTodo = false

	// search data todo
	for _, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			todoData = todo
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoData)
}

// Create CreateTodo Function here ...
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var data Todos

	json.NewDecoder(r.Body).Decode(&data)

	todos = append(todos, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// Create UpdateTodo Function here ...

// Create DeleteTodo Function here ...
