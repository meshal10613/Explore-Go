package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var ch = make(chan string)
	wg.Add(1)
	go uploadFile(ch)
	wg.Wait()

	// fileUrl := <-ch
	// fmt.Println(fileUrl)
}

func uploadFile(c chan string) {
	defer wg.Done()
	fmt.Println("uploading file....")
	time.Sleep(3 * time.Second)
	fmt.Println("file upload done!")
	fileUrl := "https://www.google.com"
	c <- fileUrl
}
