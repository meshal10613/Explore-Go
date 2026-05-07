package main

import "fmt"

func main() {
	//? Pointer in Go
	// a := 42
	// p := &a

	//* & -> address of operator
	//* * -> dereference operator

	// fmt.Println("a:", a)
	// fmt.Println("p:", p) // p is the memory address of a
	// fmt.Println("p:", *p) // *p is the value at the memory address p, which is the value of a

	// y := 10
	// change(y) // y is passed by value, so the change function receives a copy of y

	// fmt.Println("y:", y) // y remains unchanged because change function modifies a copy of y

	// z := 10
	// change2(&z) // z is passed by reference, so the change function receives a pointer to z
	// fmt.Println("z:", z) // z is changed because change2 function modifies the value at the memory address of z

	bigArray := [5]int{10, 20, 30, 40, 50}

	modifyWithoutPointer(bigArray)
	fmt.Println("After modifyWithoutPointer, bigArray:", bigArray) // bigArray remains unchanged because modifyWithoutPointer receives a copy of bigArray

	fmt.Println()

	modifyWithPointer(&bigArray)
	fmt.Println("After modifyWithPointer, bigArray:", bigArray) // bigArray is changed because modifyWithPointer receives a pointer to bigArray
}

func change(x int) {
	x = 200
	fmt.Println("x:", x)
}

func change2(x *int) {
	*x = 100
	fmt.Println("x:", *x)
}

func modifyWithoutPointer(arr [5]int) {
	arr[0] = 99
	fmt.Println("Inside without pointer,", arr)
}

func modifyWithPointer(arr *[5]int) {
	(*arr)[0] = 99
	fmt.Println("Inside with pointer,", *arr)
}
