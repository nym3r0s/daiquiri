package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GokulSrinivas/daiquiri/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	response := JsonResponse{
		Status: "404",
		Data:   "API route not found",
	}

	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		Error404(w, r)
	}
	fmt.Println(r.Form)
	var age int
	name := r.Form["name"][0]
	age, err = strconv.Atoi(r.Form["age"][0])

	// Make the new user object
	newUser := database.User{
		Name: name,
		Age:  age,
	}
	// DB operations start here

	db := database.Get_DB_Object("./database/db_config.json")

	// Check if user exists
	existingUser := new(database.User)

	db.Where(newUser).First(&existingUser)

	// fmt.Println("Existing User")
	// fmt.Println(existingUser)

	// Check if there is an exsiting user
	if existingUser.UserId == 0 {
		// No user exists
		db.Create(&newUser)
		// Set response
		response.Status = "201"
		response.Data = "User Successfully Created"
		// fmt.Println(newUser)
	} else {
		// User already exists
		// Set response
		response.Status = "400"
		response.Data = "User Already Exists!"
	}

	// Marshal JSON and send it back
	myjsonresponse, err2 := json.Marshal(response)
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	if err2 == nil {
		w.Write(myjsonresponse)
	} else {
		//peace
	}
}
