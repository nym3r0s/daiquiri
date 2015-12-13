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
	UserId int `gorm:"primary_key" valid:"-"`

	UserName  string `valid:"length(1|255)"`
	UserEmail string `sql:"UNIQUE;default:null" valid:"email"`
	UserPhone string `sql:"not null;UNIQUE" valid:"numeric,required,length(10|10)"`

	UserAge    int    `valid:"-"`
	UserAadhar string `sql:"UNIQUE;default:null" valid:"alphanum"`

	Safe bool `sql:"default:false" valid:"-"`

	PosLat  string `valid:"latitude"`
	PosLong string `valid:"longitude"`

	CreatedAt time.Time `sql:"default:now()" valid:"-"`
	UpdatedAt time.Time `sql:"default:now()" valid:"-"`
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
	// WebOtp       int
	// WebSessionId string
	// WebCreatedAt time.Time
	// WebUpdatedAt time.Time

	// App Credentials - Only one phone at a time
	AppOtp       int
	AppSessionId string
	AppCreatedAt time.Time `sql:"default:now()"`
	AppUpdatedAt time.Time `sql:"default:now()"`
}

type Admin struct {
	AdminHandle   string    `gorm:"primary_key"`
	AdminEmail    string    `sql:""`
	AdminMno      string    `sql:""`
	AdminName     string    `sql:""`
	Sudo          int       `sql:"default:0"`
	AdminPassword string    `sql:""`
	CreatedAt     time.Time `sql:"default:now()"`
	UpdatedAt     time.Time `sql:"default:now()"`
}
