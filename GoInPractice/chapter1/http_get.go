package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("http://gooogle.com")
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	res.Body.Close()
}
