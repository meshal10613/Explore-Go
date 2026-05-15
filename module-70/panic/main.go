package main

import (
	"fmt"
	"log"
	// "os"
)

func doSomething() {
	defer func() {
		fmt.Println("Defered function ran")
		r := recover()
		if r != nil {
			fmt.Println("Recoverd from", r)
		}
	}()

	panic("Something went wrong....!")
}

func doAnotherThing() {
	// fmt.Println("Done!")
	// os.Exit(1)
	log.Fatal()
}

func main() {
	//? Panic
	doSomething()

	fmt.Println("Main completed normally")
	doAnotherThing()
}
