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
	trxs := models.GetAllTransactions()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trxs)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	newTrx := &models.Transaction{}
	utils.ParseBody(r, newTrx)
	b := newTrx.CreateTransaction()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func GetTrxByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trxId := vars["id"]
	Id, err := strconv.ParseUint(trxId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	trxDetail := models.GetTransactionByID(Id)
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
	vars := mux.Vars(r)
	trxId := vars["id"]
	Id, err := strconv.ParseUint(trxId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	trxDetail := models.DeleteTransaction(Id)
	res, _ := json.Marshal(trxDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var updateTrx = &models.Transaction{}
	utils.ParseBody(r, updateTrx)

	vars := mux.Vars(r)
	trxId := vars["id"]
	Id, err := strconv.ParseUint(trxId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	trxDetail := models.UpdateTransaction(Id, *updateTrx)

	res, _ := json.Marshal(trxDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
