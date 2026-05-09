package main

import "fmt"

func main() {
	myMap := map[string]string{"Name": "John", "City": "New York", "Occupation": "Developer"}

	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}

	fmt.Print("\n")

	myArr := [3]string{"Go", "Python", "JavaScript"}
	for index, value := range myArr {
		fmt.Println(index, ":", value)
	}

	fmt.Print("\n")

	name := "Meshal"
	for index, char := range name {
		fmt.Println(index, ":", string(char))
	}
}
