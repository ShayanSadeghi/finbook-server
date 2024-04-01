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

func GetResources(w http.ResponseWriter, r *http.Request) {
	resources := models.GetAllResources()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resources)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	newResource := &models.Resource{}
	utils.ParseBody(r, newResource)
	resource := newResource.CreateResource()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}

func GetResourceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceId := vars["id"]
	Id, err := strconv.ParseUint(resourceId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	resourceDetail := models.GetResourceByID(Id)
	if resourceDetail == nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(resourceDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceId := vars["id"]
	Id, err := strconv.ParseUint(resourceId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	resourceDetail := models.DeleteResource(Id)
	res, _ := json.Marshal(resourceDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateResource(w http.ResponseWriter, r *http.Request) {
	var updateResource = &models.Resource{}
	utils.ParseBody(r, updateResource)

	vars := mux.Vars(r)
	resourceId := vars["id"]
	Id, err := strconv.ParseUint(resourceId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	resourceDetail := models.UpdateResource(Id, *updateResource)

	res, _ := json.Marshal(resourceDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
