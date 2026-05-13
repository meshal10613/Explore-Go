package main

import "fmt"

func deffered(result int) {
	fmt.Println("defered result", result)
}

func example() int {
	result := 10

	defer deffered(result)

	fmt.Println("I am from example fn:", result)
	result += 100
	return result
}

func main() {
	//? Defer Function
	// defer fmt.Println("I am defered function...")

	// fmt.Println("I am from main function...")

	example()
}
