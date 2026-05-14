package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//? Concurrency

	var start = time.Now()

	// wg.Add(1) //? Counter 1
	// go uploadFile()

	// wg.Add(1) //? Counter 2
	// go.saveTsaveToDB()

	// wg.Add(1) //? Counter 3
	// go.sendEsendEmail()

	wg.Go(uploadFile)
	wg.Go(saveToDB)
	wg.Go(sendEmail)

	//? Counter 0
	wg.Wait() //? Waiting for all goroutines

	var end = time.Since(start)
	fmt.Println(end)
}

func uploadFile() {
	// defer wg.Done() //? Counter - 1
	fmt.Println("uploading file....")
	time.Sleep(3 * time.Second)
	fmt.Println("file uploaded!")
}

func saveToDB() {
	// defer wg.Done() //? Counter - 1
	fmt.Println("saving to db....")
	time.Sleep(2 * time.Second)
	fmt.Println("saved to db!")
}

func sendEmail() {
	// defer wg.Done() //? Counter - 1
	fmt.Println("sending email....")
	time.Sleep(1 * time.Second)
	fmt.Println("email sent!")
}
