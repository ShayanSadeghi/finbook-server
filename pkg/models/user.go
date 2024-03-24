package models

import (
	"finbook-server/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	id       uint64 `gorm:"primarykey"json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	// db.Save(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserByID(Id uint64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id uint64) User {
	var user User
	db.Where("id=?", Id).Delete(user)
	return user
}
