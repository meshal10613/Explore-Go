package main

import (
	"fmt"
	"unicode/utf8"
)

func makeCofee(coffeeType string, isSuger bool) string {
	suger := "without"
	if isSuger {
		suger = "with"
	}

	return fmt.Sprintf("Making %s coffee %s suger\n", coffeeType, suger)
}

func makeName(myName string) (name string, nameChar int) {
	name = myName
	nameChar = utf8.RuneCountInString(myName)
	return
}

func main() {
	// var name string
	// name = "World"

	// n := "Go"

	// fmt.Println("Hello, " + n + " " + name)

	//? Grouped Variable Declaration
	// var (
	// 	a string = "Meshal"
	// 	b int    = 10613
	// 	c bool   = true
	// )

	// fmt.Println(a, b, c)

	//? Multiple Variable Declaration
	// var x, y int = 10, 20
	// fmt.Println(x, y)

	//?
	// const pi float32 = 3.14
	// fmt.Println(pi)

	//? Only positive number
	// const id uint = 10613

	// fmt.Println(id)

	//? Zero Value
	// var age int //* default value of int is 0
	// var name string //* default value of string is empty string ""
	// var isAdmin bool //* default value of bool is false
	// var score float64 //* default value of float64 is 0.0

	// fmt.Println(score)

	//? Printf and Sprintf
	// var name string = "Meshal"
	// var age int = 22
	// rating := 4.5
	// fmt.Printf("My name is %s & my age is %d & my rating is %.1f", name, age, rating)
	// result := fmt.Sprintf("My name is %s & my age is %d & my rating is %.1f", name, age, rating)
	// fmt.Println(result)

	//? Function
	// result := makeCofee("Espresso", true)
	// fmt.Print(result)

	//? Multiple Return Values
	// name, nameChar := makeName("Meshal")
	// fmt.Printf("My name is %s & my name has %d characters", name, nameChar)

	//? Task
	task()
}
