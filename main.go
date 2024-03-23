package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello there")
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()

	users = append(users, User{ID: 1, Username: "Shayan", Email: "shayan@mail.com", Password: "1234"})
	users = append(users, User{ID: 2, Username: "Sara", Email: "sara@mail.com", Password: "1234"})

	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("App is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
