package admincontroller

import (
	// "crypto/sha1"
	"fmt"
	// "math/rand"
	"encoding/json"
	"net/http"
	// "strconv"
	// "time"

	"github.com/GokulSrinivas/daiquiri/controllers"
	"github.com/GokulSrinivas/daiquiri/database"
	// "github.com/GokulSrinivas/daiquiri/mail"
	// "github.com/asaskevich/govalidator"
)

func GetMapData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		// fmt.Println("Error: ", err)
		controllers.WriteJson(w, r, "ERR", "Incorrect Data")
		return
	}

	admin_handle := r.FormValue("admin_handle")
	admin_pass := r.FormValue("admin_password")

	if admin_handle == "" || admin_pass == "" {
		controllers.WriteJson(w, r, "ERR", "Incorrect Data, Empty value")
		return
	}

	var newAdmin database.Admin
	// newAdmin := database.Admin{
	// AdminHandle:   admin_handle,
	// AdminPassword: admin_pass,
	// }

	db := database.Get_DB_Object("./database/db_config.json")

	db.Where("admin_handle = ?", admin_handle).Where("admin_password = ?", admin_pass).First(&newAdmin)

	fmt.Println(newAdmin.AdminHandle == "")

	if newAdmin.AdminHandle == "" {
		controllers.WriteJson(w, r, "AUTH", "Invalid Credentials")
		return
	}

	// Authenticated User. Send the data

	var users []database.User

	db.Select([]string{"pos_lat", "pos_long", "safe"}).Find(&users)
	fmt.Println(users)

	type jsonuser struct {
		Lat  string `json:"lat"`
		Long string `json:"lng"`
		Safe bool   `json:"safe"`
	}

	var returnarr []jsonuser

	for i := 0; i < len(users); i++ {
		// fmt.Println(users[i].PosLat)
		// fmt.Println(users[i].PosLong)
		// fmt.Println(users[i].Safe)
		returnarr = append(returnarr, jsonuser{
			Lat:  users[i].PosLat,
			Long: users[i].PosLong,
			Safe: users[i].Safe,
		})
	}

	fmt.Println(returnarr)

	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Status string     `json:"status"`
		Data   []jsonuser `json:"data"`
	}{
		Status: "OK",
		Data:   returnarr,
	}
	myjsonresponse, err := json.Marshal(data)

	if err == nil {
		w.Write(myjsonresponse)
	} else {
		//peace
		fmt.Println("Error in response data", err)
	}

}
