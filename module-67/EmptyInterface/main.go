package main

import "fmt"

// any == interface{}

func Process(data any) {
	stringData, ok := data.(string)
	if !ok {
		fmt.Println("Data is not string")
		return
	}

	fmt.Println(stringData)
}

func main() {
	// var data interface{}

	// data = 10
	// data = "Hello"
	// data = true
	// fmt.Println(data)

	Process("Hello")
}
