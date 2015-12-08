package main

import (
	"github.com/GokulSrinivas/daiquiri/database"
)

func main() {
	database.Config_Init("./database/db_config.json")
	database.DB_Init()
	database.Run_Migrations()
}
