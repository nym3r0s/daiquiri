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

	UserName  string
	UserEmail string `sql:"UNIQUE"`
	UserPhone string `sql:"not null;UNIQUE"`

	UserAge    int
	UserAadhar string `sql:"UNIQUE"`

	Safe bool `sql:"default:false"`

	PosLat  float64
	PosLong float64

	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time `sql:"default:now()"`
}

type Friend struct {
	FriendId      int  `gorm:"primary_key"`
	UserId        int  // The User who allows people to track him
	User          User // User Object that has Id UserId
	TrackableBy   int  // The users who can track the user (UserId)
	TrackableUser User // User Object that has Id TrackableBy
}

type AppTokens struct {
	ReqId  int  `gorm:"primary_key"`
	UserId int  `sql:"UNIQUE"` // Id of the User
	User   User // Object so that we can get it in one query

	// Web Credentials - Only one machine at a time
	WebOtp       int
	WebSessionId string
	WebCreatedAt time.Time
	WebUpdatedAt time.Time

	// App Credentials - Only one phone at a time
	AppOtp       int
	AppSessionId string
	AppCreatedAt time.Time `sql:"default:now()"`
	AppUpdatedAt time.Time `sql:"default:now()"`
}
