package main

import (
	"time"
	"fmt"
)

func send (ch chan<- string, done <-chan bool) {
	for {
		select {
		case <- done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	msg := make(chan string)
	done := make(chan bool)

	util := time.After(5 * time.Second)

	go send(msg, done)

	for {
		select {
		case m := <- msg:
			fmt.Println(m)
		case <- util:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}

}
