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
	// x := 5
	// {
	// 	fmt.Printf("Inside the block: %d\n", x)
	// }
	// fmt.Printf("Outside the block: %d\n", x)

	//? If, Else & Switch Statement
	num := 10
	// if num%2 == 0 {
	// 	fmt.Println("Even")
	// } else if num%3 == 0 {
	// 	fmt.Println("Divisible by 3")
	// } else {
	// 	fmt.Println("Odd")
	// }

	switch {
	case num > 10:
		fmt.Println("Greater than 10")
	case num < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Equal to 10")
	}
}
