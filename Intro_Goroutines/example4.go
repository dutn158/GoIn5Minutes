/**
	Buffered Channel
	it is also possible to pass a second parameter to make function when creating a channel
	this creates  a buffered channel with a capacity of 1.
 */
package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type HomePageSize struct {
	URL string
	Size int
}

func main() {
	urls := []string {
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL: url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <- results
		fmt.Println(result.URL)
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)

}
