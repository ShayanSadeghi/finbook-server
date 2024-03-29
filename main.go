package main

import (
	"finbook-server/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUsersRoutes(r)
	routes.RegisterBankRoutes(r)
	routes.RegisterAccountRoutes(r)
	http.Handle("/", r)

	fmt.Println("App is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
