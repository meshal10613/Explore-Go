package main

import (
	"fmt"
	"net/http"
	"time"
)

//? Website Health Checker

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.wrong-url.com",
	}

	startTime := time.Now()

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil || res.StatusCode != 200 {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s is up and running\n", url)
		res.Body.Close()
	}

	fmt.Println("Time taken", time.Since(startTime))
}
