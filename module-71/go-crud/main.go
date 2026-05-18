package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 5000
	http.HandleFunc("/", rootHandler)
	fmt.Printf("Server is running at http://localhost:%d", port)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server....")
}
