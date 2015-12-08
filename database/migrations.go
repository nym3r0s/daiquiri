package database

import (
	// "database/sql"
	// "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

// Structure for DB Credentials
type DbConfig struct {
	Username string
	Password string
	Database string
}

// User Table
type User struct {
	UserId int `gorm:"primary_key"`
	Name   string
	Age    int

	CreatedAt time.Time
	UpdatedAt time.Time
}
