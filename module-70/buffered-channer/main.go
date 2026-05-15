package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "File uploaded!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "File saved!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Email sent!"
	}()

	for range 3 {
		fmt.Println(<-ch)
	}
}
