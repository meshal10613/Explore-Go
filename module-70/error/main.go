package main

import "fmt"

var err error

func main() {
	//? Error handling
	data, err := divide(0, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func divide(a, b int) (int, error) {
	if b == 0 || (a == 0 && b == 0) {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
