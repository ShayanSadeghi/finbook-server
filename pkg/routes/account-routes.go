package routes

import (
	"finbook-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterAccountRoutes = func(router *mux.Router) {
	router.HandleFunc("/accounts/", controllers.GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", controllers.GetAccountByID).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", controllers.DeleteAccount).Methods("DELETE")
}
