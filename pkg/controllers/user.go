package controllers

import (
	"encoding/json"
	"finbook-server/pkg/models"
	"finbook-server/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	data := &models.User{}
	utils.ParseBody(r, data)
	token, err := models.LoginByEmail(data.Email, data.Password)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Result{Success: false, Message: err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Result{Success: true, Message: "Login Successfully", Token: token})

	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u, err := newUser.CreateUser()
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["UserId"]
	Id, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetail := models.GetUserByID(Id)
	if userDetail == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["UserId"]
	Id, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetail := models.DeleteUser(Id)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)

	vars := mux.Vars(r)
	userId := vars["UserId"]
	Id, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	userDetail := models.UpdateUser(Id, *updateUser)

	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
