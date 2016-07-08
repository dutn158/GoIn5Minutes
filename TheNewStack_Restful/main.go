package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handle "/" request
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex handle "/todos" request
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TodoIndex!")
}

// TodoShow handle "/todos/{todoId]}" request
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintf(w, "Todo show: ", todoID)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	http.HandleFunc("/", Index)
	http.HandleFunc("/todos", TodoIndex)
	http.HandleFunc("/todos/{todoId}", TodoShow)
	error := http.ListenAndServe(":3001", router)
	if error != nil {
		log.Fatal(error)
	}
}
