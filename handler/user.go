package handler

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"pharmacy-backend/auth"
	"pharmacy-backend/db"
)

//user to bind request
type user struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func SignIn(conn *sql.DB, s string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			err    error
			u      user                   //user instance to bind request body
			dbUser db.User                //user data read from db
			claims map[string]interface{} //claims to be added to JWT
			t      string                 //signed token
		)

		// Bind body to u
		if err = c.Bind(&u); err != nil {
			return returnError(c, http.StatusBadRequest, err, "Error binding request body")
		}

		// Get user from DB
		dbUser.New(conn)
		err = dbUser.Get(u.UserName)
		if err != nil {
			return returnError(c, http.StatusNotFound, err, "error retrieving user data")
		}

		// Check if password match to db
		if !auth.ComparePassword(dbUser.Password, u.Password) {
			return returnError(c, http.StatusUnauthorized, nil, "User/Pass mismatch")
		}

		// Claims to be added to JWT
		claims = make(map[string]interface{})
		claims["id"] = dbUser.Id
		claims["user"] = dbUser.Username

		// Sign token
		t, err = auth.CreateJWT(s, claims)
		if err != nil {
			return returnError(c, http.StatusInternalServerError, err, "error signing JWT")
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"token": t,
		})
	}
}
