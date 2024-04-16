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
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	accounts, err := models.GetAllAccounts(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "failed to authorize")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	newAccount := &models.Account{}

	tokenString := r.Header.Get("Authorization")
	utils.ParseBody(r, newAccount)

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}

	tokenString = tokenString[len("Bearer "):]
	a, err := newAccount.CreateAccount(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "account creation error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenString := r.Header.Get("Authorization")
	accountId := vars["id"]
	Id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	accountDetail, err := models.GetAccountByID(Id, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error on fetching data")
		return
	}

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
	tokenString := r.Header.Get("Authorization")
	vars := mux.Vars(r)

	accountId := vars["id"]
	Id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	accountDetail, err := models.DeleteAccount(Id, tokenString)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, err.Error())
		return
	}

	res, _ := json.Marshal(accountDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var UpdateAccount = &models.Account{}

	tokenString := r.Header.Get("Authorization")
	utils.ParseBody(r, UpdateAccount)

	vars := mux.Vars(r)
	accountId := vars["id"]
	Id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	accountDetail, err := models.UpdateAccount(Id, *UpdateAccount, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error on updating data")
		return
	}

	res, _ := json.Marshal(accountDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
