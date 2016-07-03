package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{}

type User struct {
	FullName  string `json:"username"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrade.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, _ := conn.ReadMessage()
				conn.WriteMessage(mType, msg)
				conn.WriteJSON(User{
					FullName:  "Hello There",
					FirstName: "Hello",
					LastName:  "There",
				})
			}
		}(conn)
	})

	err := http.ListenAndServe(":3002", nil)
	if err != nil {
		fmt.Println(err)
	}
}
