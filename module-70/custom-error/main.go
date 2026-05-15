package main

import "fmt"

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
		return "", &CustomeError{success: false, code: 401, message: "Unauthorized Access"}
	}

	return "Login Success", nil
}

func main() {
	//? Custom Error
	data, err := login("123nk4")
	if err != nil {
		fmt.Println("Success:", err.(*CustomeError).success)
		fmt.Println("Code:", err.(*CustomeError).code)
		fmt.Println("Message:", err.(*CustomeError).message)
		return
	}
	fmt.Println(data)
}
