package routes

import (
	"finbook-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBankRoutes = func(router *mux.Router) {
	router.HandleFunc("/banks/", controllers.GetBanks).Methods("GET")
	router.HandleFunc("/banks/", controllers.CreateBank).Methods("POST")
	router.HandleFunc("/banks/{id}", controllers.GetBankByID).Methods("GET")
	router.HandleFunc("/banks/{id}", controllers.UpdateBank).Methods("PUT")
	router.HandleFunc("/banks/{id}", controllers.DeleteBank).Methods("DELETE")
}
