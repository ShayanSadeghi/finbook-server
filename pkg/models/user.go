package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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
	db.QueryFields = true
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return &User{}, err
	}

	u.Password = string(hash)
	db.Create(&u)
	return u, nil
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

func LoginByEmail(email string, password string) (string, error) {
	var user User
	result := db.First(&user, User{Email: email})
	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	if result := db.First(&user, User{Email: email, Password: password}); result.Error != nil {
		fmt.Println(result.Error)
		return "", result.Error
	}

	return createToken(user)

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
