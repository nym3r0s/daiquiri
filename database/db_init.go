package database

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Config_Init(path string) {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	config := DbConfig{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	// set global var
	db_cred = config
	// Debug
	// fmt.Println(config)
}

func DB_Init() {

	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// Debug
	// fmt.Println(db_cred)

	d, err := gorm.Open("mysql", db_cred.Username+":"+db_cred.Password+"@/"+db_cred.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	// set global var
	db = d

	db.SingularTable(true)
}

func Run_Migrations() {
	err := db.CreateTable(&User{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created Table")
	}

}

func Get_DB_Object(path string) gorm.DB {
	Config_Init(path)
	DB_Init()
	return db
}
