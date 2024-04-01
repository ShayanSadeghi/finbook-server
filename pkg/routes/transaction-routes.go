package routes

import (
	"finbook-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterTrxRoutes = func(router *mux.Router) {
	router.HandleFunc("/transaction/", controllers.GetTransactions).Methods("GET")
	router.HandleFunc("/transaction/", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/transaction/{id}", controllers.GetTrxByID).Methods("GET")
	router.HandleFunc("/transaction/{id}", controllers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transaction/{id}", controllers.DeleteTransaction).Methods("DELETE")
}
