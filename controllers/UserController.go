package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GokulSrinivas/daiquiri/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		Error404(w, r)
	}
	fmt.Println(r.Form)
	var age int

	name := r.Form["name"][0]
	age, err = strconv.Atoi(r.Form["age"][0])

	if err != nil {
		WriteJson(w, r, "400", "Incorrect Data")
		return
	}
	// Make the new user object
	newUser := database.User{
		UserName: name,
		UserAge:  age,
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

		if newUser.UserId == 0 {
			WriteJson(w, r, "400", "User Already Exists")
			return
		}
		// Set response
		WriteJson(w, r, "200", strconv.Itoa(newUser.UserId))
		return
		// fmt.Println(newUser)
	} else {
		// User already exists
		// Set response
		WriteJson(w, r, "400", "User Already Exists")
		return
	}
}

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
