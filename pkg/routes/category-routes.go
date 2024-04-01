package routes

import (
	"finbook-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterResourceCatRoutes = func(router *mux.Router) {
	router.HandleFunc("/category/", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/category/", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", controllers.GetCategoryByID).Methods("GET")
	router.HandleFunc("/category/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", controllers.DeleteCategory).Methods("DELETE")
}
