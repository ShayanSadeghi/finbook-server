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

func GetBanks(w http.ResponseWriter, r *http.Request) {
	banks := models.GetAllBanks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banks)
}

func CreateBank(w http.ResponseWriter, r *http.Request) {
	newBank := &models.Bank{}
	utils.ParseBody(r, newBank)
	b := newBank.CreateBank()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func GetBankByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankId := vars["id"]
	Id, err := strconv.ParseUint(bankId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bankDetail := models.GetBankByID(Id)
	if bankDetail == nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(bankDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankId := vars["id"]
	Id, err := strconv.ParseUint(bankId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bankDetail := models.DeleteBank(Id)
	res, _ := json.Marshal(bankDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	var updateBank = &models.Bank{}
	utils.ParseBody(r, updateBank)

	vars := mux.Vars(r)
	bankId := vars["id"]
	Id, err := strconv.ParseUint(bankId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	bankDetail := models.UpdateBank(Id, *updateBank)

	res, _ := json.Marshal(bankDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
