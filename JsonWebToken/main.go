package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"net/http"
	"fmt"
	"strings"
	"time"
	"net/http"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func setToken(w http.ResponseWriter, r *http.Request) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)
	claims := MyCustomClaims{
		"username",
		jwt.StandardClaims{
			ExpiresAt:expireToken,
			Issuer:"example.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))

	cookie := http.Cookie{Name:"Auth", Value:signedToken, Expires: expireCookie, HttpOnly:true}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

func main() {

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}
