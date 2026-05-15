package main

import (
	"errors"
	"fmt"
)

var err error

func main() {
	//? Error handling
	data, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func divide(a, b int) (int, error) {
	if b == 0 || (a == 0 && b == 0) {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
