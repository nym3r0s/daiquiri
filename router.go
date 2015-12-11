package main

import (
	"net/http"

	"github.com/GokulSrinivas/daiquiri/controllers/errorcontroller"
	"github.com/GokulSrinivas/daiquiri/controllers/usercontroller"
	// "github.com/GokulSrinivas/daiquiri/mail"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()

	apirouter.HandleFunc("/user/create", usercontroller.CreateUser).Methods("POST")
	apirouter.HandleFunc("/user/mail", usercontroller.SendOTPEmail).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(errorcontroller.Error404)
	apirouter.NotFoundHandler = http.HandlerFunc(errorcontroller.Error404)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":4000")
}
