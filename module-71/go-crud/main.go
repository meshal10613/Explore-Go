package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	port := 5000
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	fmt.Printf("Server is running at http://localhost:%d", port)
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println(mux)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server....")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is healthy....")
}
