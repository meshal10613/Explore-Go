package main

import "fmt"

func main() {
	//? Array in Go
	// var num [6]int
	// num[0] = 10 // start from 0 index
	// fmt.Println("Array: ", num)
	// fmt.Println("Length of Array: ", len(num))

	// for i := 0; i < len(num); i++ {
	// 	fmt.Printf("%d index is: %d\n", i, num[i])
	// }

	// numbers := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Array:", numbers)

	//? Slice in Go
	// orders := []int{10, 20, 30, 40, 50}
	// slice1 := orders[1:4] // slice from index 1 to 3 (4 is exclusive)
	// full := orders[:]     // slice of the entire array

	// fmt.Println("Slice:", slice1)
	// fmt.Println("Full:", full)

	// fmt.Println("Length of Slice:", len(slice1)) // length is the number of elements in the slice
	// fmt.Println("Capacity of Slice:", cap(slice1)) // capacity is the number of elements in the underlying array starting from the first element of the slice

	// slice1 = append(slice1, 500)
	// slice1 = append(slice1, 600)
	// slice1 = append(slice1, 700)
	// slice1 = append(slice1, 800)
	// fmt.Println("Slice after append:", slice1)

	// fmt.Println("Length of Slice:", len(slice1))
	// fmt.Println("Capacity of Slice:", cap(slice1))

	// var slice2 []int // nil slice
	// slice2 = append(slice2, 100)
	// slice2 = append(slice2, 200)
	// slice2 = append(slice2, 300)
	// slice2 = append(slice2, 400)
	// fmt.Println("Slice 2:", slice2)

	//? Pointer in Go
	// a := 42
	// p := &a

	//* & -> address of operator
	//* * -> dereference operator

	// fmt.Println("a:", a)
	// fmt.Println("p:", p) // p is the memory address of a
	// fmt.Println("p:", *p) // *p is the value at the memory address p, which is the value of a

	y := 10
	change(y) // y is passed by value, so the change function receives a copy of y

	fmt.Println("y:", y) // y remains unchanged because change function modifies a copy of y

	z := 10
	change2(&z) // z is passed by reference, so the change function receives a pointer to z
	fmt.Println("z:", z) // z is changed because change2 function modifies the value at the memory address of z
}

func change(x int) {
	x = 200
	fmt.Println("x:", x)
}

func change2(x *int) {
	*x = 100
	fmt.Println("x:", *x)
}
