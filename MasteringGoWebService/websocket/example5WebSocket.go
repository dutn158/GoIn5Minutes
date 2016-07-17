package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"strconv"
	"net/http"
	"os"
)

func EchoLengthServer(ws *websocket.Conn) {
	var msg string
	for {
		websocket.Message.Receive(ws, &msg)
		fmt.Println("Get message", msg)
		length := len(msg)
		if err := websocket.Message.Send(ws, strconv.FormatInt(int64(length), 10)); err != nil {
			fmt.Println("Cant send message length")
			break
		}
	}
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	http.Handle("/length", websocket.Handler(EchoLengthServer))

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "MasteringGoWebService/websocket/websocket.html")
	})

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
	}
}
