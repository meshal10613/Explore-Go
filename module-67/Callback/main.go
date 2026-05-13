package main

import "fmt"

//? Callback Function
//? First Class Citizen =>

// func process(hello func()) {
// 	hello()
// }

func calculate(a int, b int, operation func(x, y int) int) int {
	result := operation(a, b)
	return result
}

func main() {
	// greet := func() {
	// 	fmt.Println("Hello")
	// }

	// process(greet)

	// add := func(x int, y int) int {
	// 	return x + y
	// }
	// multiply := func(x int, y int) int {
	// 	return x * y
	// }

	// fmt.Println(calculate(10, 20, add))
	// fmt.Println(calculate(10, 20, multiply))

	//? Anonymous Callback Function
	callback := calculate(3, 5, func(x int, y int) int {
		return x + y
	})
	fmt.Println(callback)
}
