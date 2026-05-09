package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func (d Cat) Speak() {
	fmt.Println("Meow!")
}

func makeSound(a Animal) {
	a.Speak()
}

func main() {
	//? Interface
	dexter := Dog{}
	makeSound(dexter)
}
