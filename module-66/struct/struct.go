package main

import "fmt"

// type user struct {
// 	name       string
// 	age        int
// 	isLoggedIn bool
// 	greet      func()
// }

type User struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	//? More on Structs
	// user1 := user{
	// 	name:       "John Doe",
	// 	age:        30,
	// 	isLoggedIn: false,
	// }

	// user1.greet = func() {
	// 	fmt.Println("Hello,", user1.name)
	// }
	// user1.greet()

	user2 := User{name: "John", age: 25, isLoggedIn: false}
	// user3 := User{name: "Doe", age: 30, isLoggedIn: true}

	user2.login()

	fmt.Printf("%+v", user2)
}

//? Method with Structs
func (user User) greet() {
	fmt.Println("Hello,", user.name)
}

func (user *User) login() {
	(*user).isLoggedIn = true
	fmt.Println(user.name, "is now logged in.")
}
