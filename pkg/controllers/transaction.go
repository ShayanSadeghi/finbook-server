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

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	trxs, err := models.GetAllTransactions(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trxs)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	newTrx := &models.Transaction{}
	utils.ParseBody(r, newTrx)
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	b, err := newTrx.CreateTransaction(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func GetTrxByID(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	vars := mux.Vars(r)
	trxId := vars["id"]
	Id, err := strconv.ParseUint(trxId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	trxDetail, err := models.GetTransactionByID(Id, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return
	}

	if trxDetail == nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(trxDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	vars := mux.Vars(r)
	trxId := vars["id"]
	Id, err := strconv.ParseUint(trxId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing Authorization Header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	trxDetail, err := models.DeleteTransaction(Id, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, err.Error())
		return
	}

	res, _ := json.Marshal(trxDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var updateTrx = &models.Transaction{}

	tokenString := r.Header.Get("Authorization")
	utils.ParseBody(r, updateTrx)

	vars := mux.Vars(r)
	trxId := vars["id"]
	Id, err := strconv.ParseUint(trxId, 10, 64)

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

	trxDetail, err := models.UpdateTransaction(Id, *updateTrx, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, err.Error())
		return
	}

	res, _ := json.Marshal(trxDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
