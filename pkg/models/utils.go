package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

type Env struct {
	SECRET_KEY string
}

var envFile, err = godotenv.Read(".env")

var secretKey = []byte(envFile["SECRET_KEY"])

type UserClaim struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func createToken(user User) (string, error) {
	secretKey = []byte(envFile["SECRET_KEY"])
	claims := UserClaim{user.Id, user.Email, jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return &UserClaim{}, err
	}

	if !token.Valid {
		return &UserClaim{}, fmt.Errorf("invalid token")
	}

	return token.Claims.(*UserClaim), nil
}

func verifyAccount(accountId uint64, tokenString string) (bool, error) {

	var accounts []Account

	user, err := verifyToken(tokenString)

	if err != nil {
		return false, err
	}

	db.Find(&accounts, Account{UserID: user.Id})

	for i := range len(accounts) {
		if accounts[i].Id == accountId {
			return true, nil
		}
	}

	return false, nil
}
