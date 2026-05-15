package main

import (
	"errors"
	"fmt"
)

type CustomeError struct {
	success bool
	code    int
	message string
}

func (cu *CustomeError) Error() string {
	return cu.message
}

func login(password string) (string, error) {
	if password != "1234" {
		return "",
			errors.New("Unauthorized Access")
		// &CustomeError{success: false, code: 401, message: "Unauthorized Access"}
	}

	return "Login Success", nil
}

func main() {
	//? Custom Error
	data, err := login("1234")
	if err != nil {
		// customError, ok := err.(*CustomeError)

		// if !ok {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println("Success:", customError.success)
		// 	fmt.Println("Code:", customError.code)
		// 	fmt.Println("Message:", customError.message)
		// }

		if customError, ok := err.(*CustomeError); ok {
			fmt.Println("Success:", customError.success)
			fmt.Println("Code:", customError.code)
			fmt.Println("Message:", customError.message)
		}

		return
	}
	fmt.Println(data)

	users := map[int]string{
		1: "John",
		2: "Doe",
	}
	if name, ok := users[1]; ok {
		fmt.Println(name)
	}
}
