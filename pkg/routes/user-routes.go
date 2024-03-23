package routes

import (
	"finbook-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUsersRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
}
