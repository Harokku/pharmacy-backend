package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

//HashAndSalt take a plain password and return an hashed and salted version of it
func HashAndSalt(pwd string) (string, error) {
	var (
		err  error
		hash []byte // Hashed and salted password
	)
	// Return error if passed a blank password
	if pwd == "" {
		return "", errors.New("can't compute blank password")
	}
	hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

//ComparePassword take an hashed password, a plain text password and test if they match.
func ComparePassword(hashPwd string, plainPwd string) bool {
	var (
		err error
	)
	// Return failed match if passed a blank password
	if plainPwd == "" {
		return false
	}
	// Check if passed password and stored password match
	err = bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}
