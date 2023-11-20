package sas_user

import "golang.org/x/crypto/bcrypt"

/*
Checks if username exist
*/
func checkIfUserNameExist(username string) bool {
	return false
}

/*
Hash a password
*/
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
checked if the pashword hash is correct
*/
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
Checks the quality of the password.
Returns false if the quality is not ok
*/
func checkPasswordQuality(password string) (ok bool) {
	return true
}
