package sas_user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID      string
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Role      string
	Phone     string
	EMail     string
	API       bool
	SecAuth   bool
	TotpKey   string
	db        *gorm.DB
}
