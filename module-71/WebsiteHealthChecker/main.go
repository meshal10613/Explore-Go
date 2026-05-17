package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	url    string
	status string
	error  error
}

// ? Website Health Checker
func checkWebsiteURL(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		ch <- Result{
			url:    url,
			status: "down",
			error:  err,
		}
		return
	}
	ch <- Result{
		url:    url,
		status: "is up & running",
		error:  nil,
	}

	defer res.Body.Close()
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.wrong-url.com",
		"https://www.web.programming-hero.com",
	}

	ch := make(chan Result)

	startTime := time.Now()

	for _, url := range urls {
		go checkWebsiteURL(url, ch)
	}

	for range urls {
		result := <-ch
		fmt.Println(result.url, "is", result.status)
	}

	fmt.Println("\nTime taken", time.Since(startTime))
}
