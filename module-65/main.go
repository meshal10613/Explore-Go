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
	orders := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice1 := orders[1:4] // slice from index 1 to 3 (4 is exclusive)
	full := orders[:]     // slice of the entire array

	fmt.Println("Slice:", slice1)
	fmt.Println("Full:", full)

	fmt.Println("Length of Slice:", len(slice1))
	fmt.Println("Capacity of Slice:", cap(slice1))

	slice1 = append(slice1, 500)
	fmt.Println("Slice after append:", slice1)

	fmt.Println("Length of Slice:", len(slice1))
	fmt.Println("Capacity of Slice:", cap(slice1))
}
