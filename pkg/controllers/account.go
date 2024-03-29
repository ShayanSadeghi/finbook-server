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

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := models.GetAllAccounts()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	newAccount := &models.Account{}
	utils.ParseBody(r, newAccount)
	a := newAccount.CreateAccount()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]
	Id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	accountDetail := models.GetAccountByID(Id)
	if accountDetail == nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(accountDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]
	Id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	accountDetail := models.DeleteAccount(Id)
	res, _ := json.Marshal(accountDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var UpdateAccount = &models.Account{}
	utils.ParseBody(r, UpdateAccount)

	vars := mux.Vars(r)
	accountId := vars["id"]
	Id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	accountDetail := models.UpdateAccount(Id, *UpdateAccount)

	res, _ := json.Marshal(accountDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
