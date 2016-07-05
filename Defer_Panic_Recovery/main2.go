package main

import (
	"fmt"
)

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovery from f ", r)
		}
	}()
	g(0)
}

func g(i int) {
	if i > 3 {
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in g ", i)
	fmt.Println(i)
	g(i + 1)
}

func main() {
	f()
}
