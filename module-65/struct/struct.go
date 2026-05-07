package main

import "fmt"

type User struct {
	username string
	email    string
	info     additionalInfo
}

type additionalInfo struct {
	phone   int
	address string
}

type User2 struct {
	name  string
	email string
	age   int
}

func main() {
	//? Struct in Go
	// user := User {"John Doe", "john@doe.com"} // position based initialization
	// user := User{
	// 	username: "John Doe",
	// 	email:    "john@doe.com",
	// } // key based initialization

	// user.email = "john.doe@gmail.com"

	// fmt.Println(user)
	// fmt.Printf("%+v \n", user) // %+v will print the field names along with their values
	// fmt.Println(user.email)

	// var user1 User
	// user1.username = "Lorem Ipsum"
	// user1.email = "lorem@ipsum.com"

	// fmt.Println(user1)

	// user := User{
	// 	username: "John Doe",
	// 	email:    "john@doe.com",
	// 	info: additionalInfo{
	// 		phone:   1234567890,
	// 		address: "123 Main Street",
	// 	},
	// }

	// fmt.Printf("%+v", user)

	//? Constructor Function
	newUser2 := func(name string, email string, age int) User2 {
		if name == "" || email == "" || age <= 0 {
			fmt.Println("Invalid input")
			return User2{}
		}

		return User2{
			name:  name,
			email: email,
			age:   age,
		}
	}

	Meshal := newUser2("Meshal", "meshal@example.com", 0)
	fmt.Printf("%+v", Meshal)
}
