package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	URL string
	Dur time.Duration
}

func getter(url string, ch chan<- result) {
	defer close(ch)
	start := time.Now()
	if _, err := http.Get(url); err != nil {
		ch <- result{URL: url, Dur: time.Duration(-1)}
		return
	}
	ch <- result{URL: url, Dur: time.Now().Sub(start)}
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://bing.com", "https://duckduckgo.com", "https://myweb-javanewbie.rhcloud.com"}
	durCh := make(chan result)

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		_ = "breakpoint"
		ch := make(chan result)
		go getter(url, ch)

		go func() {
			defer wg.Done()
			t := <-ch
			durCh <- t
		}()
	}

	go func() {
		wg.Wait()
		close(durCh)
	}()

	for t := range durCh {
		_ = "breakpoint"
		fmt.Println(t)
	}

}
