package usercontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GokulSrinivas/daiquiri/controllers"
	"github.com/GokulSrinivas/daiquiri/database"
	"github.com/asaskevich/govalidator"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}
	fmt.Println(r.Form)

	// Getting form data
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	age, ageerr := strconv.Atoi(r.FormValue("age"))
	lat := r.FormValue("lat")
	long := r.FormValue("long")

	aadhar := r.FormValue("aadhar")

	// Empty Field Check (Required)
	emptyerr := false

	emptyerr = emptyerr || name == ""
	emptyerr = emptyerr || email == ""
	emptyerr = emptyerr || phone == ""

	emptyerr = emptyerr || ageerr != nil

	if emptyerr {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, Missing Fields")
		return
		fmt.Println("Failed Empty check")
	}

	// Make the new user object
	newUser := database.User{
		UserName:  name,
		UserEmail: email,
		UserPhone: phone,

		UserAge: age,
		PosLat:  lat,
		PosLong: long,

		Safe: false,
	}

	if aadhar != "" {
		newUser.UserAadhar = aadhar
	}

	// Validate with Govalidator - Should catch all the errors
	_, structerr := govalidator.ValidateStruct(newUser)
	if structerr != nil {
		fmt.Println(structerr)
		controllers.WriteJson(w, r, "ERR", structerr.Error())
		return
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
			controllers.WriteJson(w, r, "ERR", "Bad Data - User Already Exists")
			return
		}
		// Set response
		controllers.WriteJson(w, r, "OK", strconv.Itoa(newUser.UserId))
		return
		// fmt.Println(newUser)
	} else {
		// User already exists
		// Set response
		controllers.WriteJson(w, r, "ERR", "User Already Exists")
		return
	}
}
