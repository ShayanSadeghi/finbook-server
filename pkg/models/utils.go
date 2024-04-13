package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var db *gorm.DB

var secretKey = []byte("my_key")

type UserClaim struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func createToken(user User) (string, error) {
	claims := UserClaim{user.Id, user.Email, jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.Id, "exp": time.Now().Add(time.Hour * 24).Unix()})
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
