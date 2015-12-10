package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

// Function to Write the response as JSON
func WriteJson(w http.ResponseWriter, r *http.Request, status string, data string) {

	w.Header().Set("Content-Type", "application/json")

	response := JsonResponse{
		Status: status,
		Data:   data,
	}

	myjsonresponse, err := json.Marshal(response)

	if err == nil {
		w.Write(myjsonresponse)
	} else {
		//peace
		fmt.Println("Error in response data", err)
	}

}
