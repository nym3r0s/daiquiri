package controllers

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func Error404(w http.ResponseWriter, r *http.Request) {
	response := JsonResponse{
		Status: "ERR",
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

func Error401(w http.ResponseWriter, r *http.Request) {
	response := JsonResponse{
		Status: "AUTH",
		Data:   "No Auth",
	}
	w.Header().Set("Content-Type", "application/json")

	myjsonresponse, err := json.Marshal(response)

	if err == nil {
		w.Write(myjsonresponse)
	} else {
		//peace
	}
}
