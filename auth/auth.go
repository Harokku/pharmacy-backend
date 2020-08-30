package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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

//CreateJWT sign a new JWT token with passed claims
//
//s string: secret key to sign the token
//
//c: claims to add to the token
func CreateJWT(s string, c map[string]interface{}) (string, error) {
	var (
		err        error
		token      *jwt.Token    //token object
		claims     jwt.MapClaims //token claims
		expiration time.Duration //token expiration
		t          string        //signed JWT token
	)
	// Parse expiration from env variable, if err set default 24H
	expiration, err = time.ParseDuration(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		expiration, _ = time.ParseDuration("24h")
	}

	token = jwt.New(jwt.SigningMethodHS256)

	// Create claims
	claims = token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(expiration)
	for k, v := range c {
		claims[k] = v
	}

	// Sign token
	t, err = token.SignedString([]byte(s))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error signing token: %v\n", err))
	}
	return t, nil
}
