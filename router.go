package main

import (
	"net/http"

	"github.com/GokulSrinivas/daiquiri/controllers/admincontroller"
	"github.com/GokulSrinivas/daiquiri/controllers/errorcontroller"
	"github.com/GokulSrinivas/daiquiri/controllers/usercontroller"
	"github.com/GokulSrinivas/daiquiri/middleware"
	// "github.com/GokulSrinivas/daiquiri/mail"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()
	adminapirouter := router.PathPrefix("/api/admin").Subrouter()
	secureapirouter := router.PathPrefix("/api/s").Subrouter()

	apirouter.HandleFunc("/user/create", usercontroller.CreateUser).Methods("POST")
	apirouter.HandleFunc("/user/auth", usercontroller.AuthOTP).Methods("POST")
	apirouter.HandleFunc("/user/login", usercontroller.Login).Methods("POST")

	secureapirouter.HandleFunc("/user/updateprofile", middleware.UserAuth(usercontroller.UpdateProfile))
	secureapirouter.HandleFunc("/user/updatestatus", middleware.UserAuth(usercontroller.UpdateStatus))

	adminapirouter.HandleFunc("/getmapdata", admincontroller.GetMapData).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(errorcontroller.Error404)
	apirouter.NotFoundHandler = http.HandlerFunc(errorcontroller.Error404)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":4000")
}
