package sas_user

import (
	"errors"
	"github.com/google/uuid"
)

/*
Discription: Creates a new User
Input:
Output:
Errors:
1)
2) "Username already exist"
*/
func (p *User) AddAUser(firstName, lastName, userName, password, phone, email string) error {
	/*
		Check if a field is empty
	*/

	//Check if username exist
	if checkIfUserNameExist(userName) {
		err := errors.New("Username already exist")
		return err
	}

	//Create UUID
	id, err := uuid.NewUUID()
	if err != nil {
		err = errors.New("Error, creating uuid: " + err.Error())
		return err
	}
	//Check Password
	if !checkPasswordQuality(password) {
		err = errors.New("Password has a bad quality")
		return err
	}

	//Hash Password
	hash, err := hashPassword(password)
	if err != nil {
		err = errors.New("Problem creating hash " + err.Error())
		return err
	}

	x := new(User)
	x.UUID = id.String()
	x.FirstName = firstName
	x.LastName = lastName
	x.UserName = userName
	x.Password = hash
	x.Phone = phone
	x.EMail = email

	err = p.db.Save(x).Error

	if err != nil {
		return err
	}

	return nil
}
