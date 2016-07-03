package main

import (
	"fmt"
	"log"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "../../public/index.html")
	fmt.Fprintf(w, "Hello World")
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", abc)
	err := http.ListenAndServeTLS(":3000", "../../cert.pem", "../../key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
