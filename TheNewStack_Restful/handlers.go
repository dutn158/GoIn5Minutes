package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handle "/" request
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex handle "/todos" request
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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

// TotoCreate... handle "/todos" request with POST
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json;chartset=UTF-8")
		w.WriteHeader(422)
		if error := json.NewEncoder(w).Encode(err); error != nil {
			panic(error)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json;chartset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
