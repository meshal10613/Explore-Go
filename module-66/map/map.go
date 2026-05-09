package main

import "fmt"

type User struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	//? Map in Go
	// myMap := make(map[string]int)

	// myMap["userPhone"] = 10
	// fmt.Println(myMap["userPhone"])

	// myMap := map[string]string{"Name": "John", "City": "New York"}
	// fmt.Println(myMap)

	myMap := map[string]User{
		"data1": {name: "John", age: 25, isLoggedIn: false},
		"data2":  {name: "Doe", age: 30, isLoggedIn: true},
	}
	fmt.Println(myMap)
}
