package usercontroller

import (
	// "crypto/sha1"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/GokulSrinivas/daiquiri/controllers"
	"github.com/GokulSrinivas/daiquiri/database"
	"github.com/GokulSrinivas/daiquiri/mail"
	"github.com/asaskevich/govalidator"
)

// Helper function to generate a random string
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// The User Related functions come here

func CreateUser(w http.ResponseWriter, r *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	myrand := rand.New(s1)
	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}
	// fmt.Println(r.Form)

	// Getting form data
	// name := r.FormValue("name")
	// email := r.FormValue("email")
	phone := r.FormValue("phone")

	// age, ageerr := strconv.Atoi(r.FormValue("age"))
	lat := r.FormValue("lat")
	long := r.FormValue("long")

	aadhar := r.FormValue("aadhar")

	// Empty Field Check (Required)
	emptyerr := false

	// emptyerr = emptyerr || name == ""
	// emptyerr = emptyerr || email == ""
	emptyerr = emptyerr || phone == ""

	// emptyerr = emptyerr || ageerr != nil

	if emptyerr {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, Missing Fields")
		return
		fmt.Println("Failed Empty check")
	}

	// Make the new user object
	newUser := database.User{
		// UserName:  name,
		// UserEmail: email,
		UserPhone: phone,

		// UserAge: age,
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

	// var dupUser database.User
	var dupAdharUser database.User
	dupAdharUser.UserId = 0

	// db.Where("user_email = ?", newUser.UserEmail).Or("user_phone = ?", newUser.UserPhone).First(&dupUser)

	if aadhar != "" {
		db.Where("user_aadhar = ?", aadhar).First(&dupAdharUser)
	}
	// Check if there is an existing user
	if existingUser.UserId == 0 && dupAdharUser.UserId == 0 {
		// No user exists
		db.Create(&newUser)

		if newUser.UserId == 0 {
			controllers.WriteJson(w, r, "EXISTS", "User Already Exists")
			return
		}
		// Set OTP
		apikey := database.AppTokens{
			UserId: newUser.UserId,
			User:   newUser,
			AppOtp: myrand.Intn(100000000),
			// AppSessionId: RandString(32),
		}

		db.Create(&apikey)
		// Send OTP
		go func() {
			SendOTPEmail(w, r, newUser.UserEmail, strconv.Itoa(apikey.AppOtp))
		}()
		// Set response
		controllers.WriteJson(w, r, "OK", "USER CREATED")
		return
		// fmt.Println(newUser)
	} else {
		// User already exists
		// Set response
		controllers.WriteJson(w, r, "EXISTS", "User Already Exists")
		return
	}
}

func SendOTPEmail(w http.ResponseWriter, r *http.Request, to string, msg string) {
	mail.Config_Init("./mail/mail_config.json")
	mail.SendMail(to, "<html>Hi. Your OTP is "+msg+"</html>")
}

func AuthOTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}

	phone := r.FormValue("phone")
	otp, err2 := strconv.Atoi(r.FormValue("otp"))

	if phone == "" || err2 != nil || r.FormValue("otp") == "" {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}

	db := database.Get_DB_Object("./database/db_config.json")
	var UserObj database.User
	var Apikey database.AppTokens
	db.Where("user_phone = ?", phone).First(&UserObj)

	db.Where("user_id = ?", UserObj.UserId).First(&Apikey)

	if UserObj.UserId == 0 || Apikey.ReqId == 0 {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}

	if otp == Apikey.AppOtp {
		Apikey.AppSessionId = RandString(32)
		db.Save(&Apikey)
		controllers.WriteJson(w, r, "OK", Apikey.AppSessionId)
		return
	} else {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Entering Update Profile")
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
	// lat := r.FormValue("lat")
	// long := r.FormValue("long")

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

	// DB operations start here
	db := database.Get_DB_Object("./database/db_config.json")
	var newUser database.User
	db.Where("user_phone = ?", phone).First(&newUser)
	// Make the new user object
	if newUser.UserId == 0 {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, Missing Fields")
		return
		fmt.Println("Failed Empty check")
	}

	newUser.UserName = name
	newUser.UserEmail = email
	newUser.UserPhone = phone
	newUser.UserAge = age

	db.Save(&newUser)

	if aadhar != "" {

		db.Exec("UPDATE user SET user_aadhar=? WHERE user_id = ?", aadhar, newUser.UserId)
	} else {
		fmt.Println("Empty Aadhar Number")
		db.Exec("UPDATE user SET user_aadhar=NULL WHERE user_id = ? and user_aadhar='' ", newUser.UserId)
	}

	controllers.WriteJson(w, r, "OK", "Updated Successfully")
	return
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data. FormParseError")
		return
	}
	fmt.Println(r.Form)

	// Getting form data
	phone := r.FormValue("phone")

	lat := r.FormValue("lat")
	long := r.FormValue("long")

	safe := r.FormValue("safe")

	var err2 error
	err2 = nil
	var safebool bool

	if safe == "" {
		safe = ""
	} else {
		safebool, err2 = strconv.ParseBool(safe)
	}

	if err2 != nil || lat == "" || long == "" || !govalidator.IsLatitude(lat) || !govalidator.IsLongitude(long) {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, Missing Fields")
		return
	}

	// DB operations start here
	db := database.Get_DB_Object("./database/db_config.json")
	var newUser database.User
	db.Where("user_phone = ?", phone).First(&newUser)
	// Make the new user object
	if newUser.UserId == 0 {
		fmt.Println(newUser)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, User not found")
		return
	}

	if safe == "" {
		safebool = newUser.Safe
	}
	newUser.PosLat = lat
	newUser.PosLong = long
	newUser.Safe = safebool
	db.Save(&newUser)

	if newUser.UserAadhar != "" {
		// Peace
		// db.Exec("UPDATE user SET user_aadhar=? WHERE user_id = ?", aadhar, newUser.UserId)
	} else {
		fmt.Println("Empty Aadhar Number")
		db.Exec("UPDATE user SET user_aadhar=NULL WHERE user_id = ? and user_aadhar='' ", newUser.UserId)
	}

	controllers.WriteJson(w, r, "OK", "Updated Successfully")
	return

}

func Login(w http.ResponseWriter, r *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	myrand := rand.New(s1)
	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data. FormParseError")
		return
	}
	fmt.Println(r.Form)

	// Getting form data
	phone := r.FormValue("phone")
	if phone == "" {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data. Missing Field")
		return
	}

	// DB operations start here
	db := database.Get_DB_Object("./database/db_config.json")
	var newUser database.User
	db.Where("user_phone = ?", phone).First(&newUser)
	// Make the new user object
	if newUser.UserId == 0 {
		fmt.Println(newUser)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, User not found")
		return
	}

	var apikey database.AppTokens

	db.Where("user_id = ?", newUser.UserId).First(&apikey)

	apikey.AppOtp = myrand.Intn(100000000)
	// apikey.AppSessionId = RandString(32)

	db.Save(&apikey)

	if apikey.ReqId == 0 {
		controllers.WriteJson(w, r, "ERR", "Err in saving new apikey")
		return
	}

	// Send OTP
	go func() {
		SendOTPEmail(w, r, newUser.UserEmail, strconv.Itoa(apikey.AppOtp))
	}()
	// Set response
	controllers.WriteJson(w, r, "OK", "USER OTP reset")
	return

}

func UpdateStatusAadhar(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data. FormParseError")
		return
	}
	fmt.Println(r.Form)

	// Getting form data
	aadhar := r.FormValue("aadhar")

	if aadhar == "" {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data. empty fields")
		return
	}

	db := database.Get_DB_Object("./database/db_config.json")

	var newUser database.User
	db.Where("user_aadhar = ?", aadhar).First(&newUser)

	if newUser.UserId != 0 {
		newUser.Safe = true
		db.Save(&newUser)
		controllers.WriteJson(w, r, "OK", "Updated Successfully")
		return
	} else {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data. User Not Found")
		return
	}
}
