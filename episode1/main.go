package main

import (
	"log"
	"net/http"

	"github.com/dutn158/GoIn5Minutes/episode1/handlers"
	"github.com/dutn158/GoIn5Minutes/episode1/storage"
)

func main() {
	db := storage.NewInMemoryDB()

	mux := http.NewServeMux()

	mux.Handle("/get", handlers.GetKey(db))

	mux.Handle("/set", handlers.PutKey(db))

	log.Printf("Serve on port 3000")

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}
