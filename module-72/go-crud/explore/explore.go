package main

import "fmt"

var users = []string{
	"Meshal",
	"Rupom",
	"Muhib",
	"Faysal",
}

func main() {
	for i, user := range users {
		if i == 2 {
			users = append(users[:i], users[i+1:]...)
			fmt.Println(user)
		}
	}
	fmt.Println(users)
}

//? users[:i] => ["Muhib"]
//? users[i + 1:]... => ["Faysal"]
//? append(["Muhib"], ["Faysal"])