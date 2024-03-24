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

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["UserId"]
	Id, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetail, _ := models.GetUserByID(Id)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Contetn-Type", "pkglibcation/json")
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
	w.Header().Set("Contetn-Type", "pkglibcation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
