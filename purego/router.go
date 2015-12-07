package main

import (
	"encoding/json"
	// "fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(Error404)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}

type JsonResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func Error404(w http.ResponseWriter, r *http.Request) {
	response := JsonResponse{
		Status: "404",
		Data:   "API route not found",
	}
	w.Header().Set("Content-Type", "application/json")

	myjsonresponse, err := json.Marshal(response)

	if err == nil {
		w.Write(myjsonresponse)
	} else {
		//peace
	}
}
