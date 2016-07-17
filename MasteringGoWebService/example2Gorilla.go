package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

type API struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]
	HelloMessage := "Hello, " + name
	message := API{
		Message: HelloMessage,
	}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	fmt.Fprintf(w, string(output))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/{user:[0-9a-zA-Z]+}", Hello)
	http.Handle("/", mux)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
