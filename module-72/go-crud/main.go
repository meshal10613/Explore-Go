package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{
		Id:    1,
		Name:  "John Doe",
		Age:   30,
		Email: "john@doe",
	},
	{
		Id:    2,
		Name:  "Jane Doe",
		Age:   26,
		Email: "jane@doe",
	},
}

func GetPathInt(r *http.Request, w http.ResponseWriter, param string) (int, bool) {
	value := r.PathValue(param)

	id, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, "Invalid "+param, http.StatusBadRequest)
		return 0, false
	}

	return id, true
}

func main() {
	mux := http.NewServeMux()

	port := 5000
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /users/{id}", getUserByIdHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserByIdHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserByIdHandler)

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

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// users, err := json.Marshal(user) //? First saves in memory
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(users)

	encoder := json.NewEncoder(w)
	encoder.Encode(users)
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := GetPathInt(r, w, "id")
	if !ok {
		return
	}

	for _, user := range users {
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "User not found", http.StatusNotFound)
}

func updateUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := GetPathInt(r, w, "id")
	if !ok {
		return
	}

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for idx, user := range users {
		if user.Id == id {
			updatedUser.Id = id
			users[idx] = updatedUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := GetPathInt(r, w, "id")
	if !ok {
		return
	}

	for i, user := range users {
		if user.Id == id {

			// delete user from slice
			// users = append(users[:i], users[i+1:]...)

			//? simple way with package
			users = slices.Delete(users, i, i+1)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "User deleted successfully",
			})
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
