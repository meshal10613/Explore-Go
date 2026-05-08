package main

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
	user2.greet() 
}

//? Method with Structs
func (user User) greet() {
	println("Hello,", user.name)
}
