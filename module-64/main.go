package main

import "fmt"

func main() {
	//? Anonymous Function
	// var makeCoffee = func() {
	// 	fmt.Printf("Making Coffee...")
	// }
	// makeCoffee()

	//? IIFE(Immediately Invoked Function Expression) Functions
	// (func() {
	// 	fmt.Printf("Making Coffee...")
	// })()

	// (func(coffeeType string) {
	// 	fmt.Printf("Making %s Coffee...", coffeeType)
	// })("Black Coffee")

	//? Variable Scope
	x := 5
	{
		fmt.Printf("Inside the block: %d\n", x)
	}
	fmt.Printf("Outside the block: %d\n", x)
}
