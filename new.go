package sas_user

import "gorm.io/gorm"

/*
Inits the porocess for a new user
*/
func NewUserManagagement(db *gorm.DB) *User {
	p := new(User)
	p.db = db
	return p
}
