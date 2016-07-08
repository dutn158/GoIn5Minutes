package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Todo struct of the Todo
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos slices
type Todos []Todo

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

	json.NewEncoder(w).Encode(todos)
}

// TodoShow handle "/todos/{todoId]}" request
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	todoID := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %q", todoID)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	error := http.ListenAndServe(":8080", router)
	if error != nil {
		log.Fatal(error)
	}
}
