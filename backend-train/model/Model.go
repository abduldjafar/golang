package model

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

// structur of job board table
type Accounts struct {
	Email      string
	Password   string
	Fullname   string
	Birthday   string
	Gender     string
	Country    string
	Plan       string
	Newsletter string
	CreatedAt time.Time
}

type Token struct {
	Email string
	*jwt.StandardClaims
}

type Exception struct {
	Message string `json:"message"`
}

type Movies struct {
	NetflixId  string
	Title      string
	Image      string
	Synopsis   string
	Type       string
	Released   string
	Runtime    string
	LargeImage string
	Unogsdate  string
	Imbid      string
	Download   string
	Rating     string
}

type Refresh struct {
	email  string
	counts string
}
