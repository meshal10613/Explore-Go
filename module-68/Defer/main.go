package main

import "fmt"

func deffered(result *int) {
	fmt.Println("defered result 1", *result)
}

func example() int {
	result := 10

	defer deffered(&result) //? If pass pointer then 110, otherwise 10
	defer func() {
		result += 50
		fmt.Println("defered result 2", result)
	}()

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
