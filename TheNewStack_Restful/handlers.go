package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handle "/" request
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex handle "/todos" request
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		panic(err)
	}

}

// TodoShow handle "/todos/{todoId]}" request
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	todoID := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %q", todoID)
}
