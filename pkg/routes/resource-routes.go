package routes

import (
	"finbook-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterResourceRoutes = func(router *mux.Router) {
	router.HandleFunc("/resource/", controllers.GetResources).Methods("GET")
	router.HandleFunc("/resource/", controllers.CreateResource).Methods("POST")
	router.HandleFunc("/resource/{id}", controllers.GetResourceByID).Methods("GET")
	router.HandleFunc("/resource/{id}", controllers.UpdateResource).Methods("PUT")
	router.HandleFunc("/resource/{id}", controllers.DeleteResource).Methods("DELETE")
}
