package main

import "time"

func main() {
	ch := make(chan string, 3)

	go func() {
		time.Sleep(3 * time.Second)
	}()

	go func() {
		time.Sleep(2 * time.Second)
	}()

	go func() {
		time.Sleep(1 * time.Second)
	}()
}
