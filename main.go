package main

// import fmt, net/http, github.com/gorilla/mux here ...
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Create main function to show "hello world" here ...

func main() {
	fmt.Println("Hello world")

	route := mux.NewRouter()
	// r.GET("/", (req, res)=> { })
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	http.ListenAndServe("localhost:5000", route)
}
