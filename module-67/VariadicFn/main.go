package main

import "fmt"

//? Variadic Function

func add(x ...int) int {
	total := 0
	for _, x := range x {
		total += x
	}

	return total
}

func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Println("💐",prefix, name)
	}
}

func main() {
	sum := add(10, 20, 30, 40)
	fmt.Println(sum)

	greet("Hello", "Nahid", "Hasnat", "Asif", "Nasir")
}
