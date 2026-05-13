package main

import "fmt"

func calculate(a int, b int, operation func(int, int) int) int {
	result := operation(a, b)
	return result
}

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	//? Higher Order Function
	// add := func(x int, y int) int {
	// 	return x + y
	// }

	// fmt.Println(calculate(10, 20, add))

	double := multiplyBy(2)
	fmt.Println(double(10))
}
