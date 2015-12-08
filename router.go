package main

import (
	"net/http"

	"github.com/GokulSrinivas/daiquiri/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()

	apirouter.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(controllers.Error404)
	apirouter.NotFoundHandler = http.HandlerFunc(controllers.Error404)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
