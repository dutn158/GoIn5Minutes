package main

import (
	"net/http"
	"regexp"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		message := "You have triggered nothing"
		testMatch, _ := regexp.MatchString("/testing[0-9]{3}", path)
		if testMatch {
			message = "You hit the Test!"
		}
		fmt.Fprintf(w, message)
	})

	http.ListenAndServe(":8080", nil)

}
