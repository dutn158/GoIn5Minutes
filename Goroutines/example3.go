/**
	Go has a special statement called select that works like a switch but for channels:
 */
package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from c1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "from c2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println("msg1:", msg1)
			case msg2 := <- c2:
				fmt.Println("msg2 ", msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			//default:
				//fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
