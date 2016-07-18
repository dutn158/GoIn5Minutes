package main

import "fmt"

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	//Squares
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	fmt.Println(<-squares)

}
