package main

import "fmt"

func coffeeOrder(customerName string, coffeeType string, price int) string {
	return fmt.Sprintf("Order for %s: %s coffee costs %d taka ☕", customerName, coffeeType, price)
}

func task() {
	// variables (customer info)
	name := "Mizan"
	coffee := "Latte"
	price := 150

	// function call
	order := coffeeOrder(name, coffee, price)

	// print result
	fmt.Println(order)

	// extra test example
	fmt.Println(coffeeOrder("Meshal", "Cappuccino", 180))
}
