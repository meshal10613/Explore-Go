package main

import "fmt"

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func makeCounter() func() int {
	count := 0

	inner := func() int {
		count++
		return count
	}

	return inner
}

func main() {
	//? Closure Function
	//? First Class Citizen =>
	// double := multiplyBy(2)
	// fmt.Println(double(10))

	next := makeCounter()

	fmt.Println(next()) //* 1
	fmt.Println(next()) //* 2
	fmt.Println(next()) //* 3
}
