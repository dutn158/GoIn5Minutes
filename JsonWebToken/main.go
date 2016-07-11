package main

import (
	"github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/context"
	//"net/http"
	//"fmt"
	//"strings"
	//"time"
	"net/http"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

func main() {

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}
