package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func (u User) CreateUser() *User {
	users = append(users, u)
	return &u
}

func GetAllUsers() []User {
	return users
}
