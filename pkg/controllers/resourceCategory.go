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

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := models.GetAllResourceCategories()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	newCategory := &models.ResourceCategory{}
	utils.ParseBody(r, newCategory)
	c := newCategory.CreateResourceCategory()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catId := vars["id"]
	Id, err := strconv.ParseUint(catId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	catDetail := models.GetResourceCategoryByID(Id)
	if catDetail == nil {
		http.Error(w, "Resource category not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(catDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catId := vars["id"]
	Id, err := strconv.ParseUint(catId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	catDetail := models.DeleteResourceCategory(Id)
	res, _ := json.Marshal(catDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updateCategory = &models.ResourceCategory{}
	utils.ParseBody(r, UpdateCategory)

	vars := mux.Vars(r)
	catId := vars["id"]
	Id, err := strconv.ParseUint(catId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	catDetail := models.UpdateResourceCategory(Id, *updateCategory)

	res, _ := json.Marshal(catDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
