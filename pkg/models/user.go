package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

func GetUserByID(Id uint64) *User {
	var getUser User
	if result := db.First(&getUser, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return &getUser
}

func DeleteUser(Id uint64) User {
	var user User
	db.Where("id=?", Id).Delete(&user)
	return user
}

func UpdateUser(Id uint64, updateUser User) User {
	userDetail := GetUserByID(Id)

	if updateUser.Username != "" {
		userDetail.Username = updateUser.Username
	}

	if updateUser.Email != "" {
		userDetail.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		userDetail.Password = updateUser.Password
	}

	db.Save(userDetail)
	return *userDetail
}
