package controllers

import (
	"encoding/json"
	"finbook-server/pkg/models"
	"finbook-server/pkg/utils"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u := newUser.CreateUser()
	res, _ := json.Marshal(u)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
