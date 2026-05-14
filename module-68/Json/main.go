package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"personName"` // Rename this field
	Age  int    `json:"-"`          // Exclude this field
	City string `json:"city"`
}

func main() {
	p := Person{Name: "John", Age: 30, City: "New York"}
	// Convert Person to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	var p2 Person
	jsonText := `{"personName":"John","city":"New York"}`
	err = json.Unmarshal([]byte(jsonText), &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
}
