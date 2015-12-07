package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/GokulSrinivas/daiquiri/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(controllers.Error404)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
