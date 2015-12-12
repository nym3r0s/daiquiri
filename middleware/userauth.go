package middleware

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/GokulSrinivas/daiquiri/controllers"
	"github.com/GokulSrinivas/daiquiri/database"
)

func UserAuth(next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()

		fmt.Println("Entering UserAuth Middleware")

		if err != nil {
			fmt.Println("Error: ", err)
			controllers.WriteJson(w, r, "ERR", "Incorrect Data")
			return
		}
		fmt.Println(r.Form)
		phone := r.FormValue("phone")
		session := r.FormValue("session")

		if phone == "" || session == "" {
			controllers.WriteJson(w, r, "ERR", "Incorrect Data Empty Field")
			return
		}

		db := database.Get_DB_Object("./database/db_config.json")
		var UserObj database.User
		var Apikey database.AppTokens
		db.Where("user_phone = ?", phone).First(&UserObj)

		db.Where("app_session_id = ?", session).First(&Apikey)

		if UserObj.UserId == 0 || Apikey.ReqId == 0 {
			controllers.WriteJson(w, r, "AUTH", "Invalid Credentials")
			return
		}

		if UserObj.UserId == Apikey.UserId {
			fmt.Println("Successfully passed UserAuth")
			next(w, r)
			return
		} else {
			controllers.WriteJson(w, r, "AUTH", "Invalid Credentials")
			return
		}

	}
}
