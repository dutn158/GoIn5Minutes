package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	formated := t.Format(time.RFC1123)
	fmt.Println(formated)
}
