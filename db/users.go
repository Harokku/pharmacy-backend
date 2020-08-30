package db

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	conn     *sql.DB //db connection
	Id       string  `json:"id"`       //user id
	Username string  `json:"username"` //username
	Password string  `json:"password"` //user's password
	Mail     string  `json:"mail"`     //user's mail
}

//New initialize connection
func (u *User) New(c *sql.DB) {
	u.conn = c
}

//GetAll retrieve all users from DB
func (u User) GetAll(dest *[]User) error {
	var (
		err  error
		rows *sql.Rows
	)
	sqlStatement := `select id, username, mail
						from users
						order by username asc 
						`
	rows, err = u.conn.Query(sqlStatement)
	if err != nil {
		return errors.New(fmt.Sprintf("error retrievinf users: %v\n", err))
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Username, &u.Mail)
		if err != nil {
			return errors.New(fmt.Sprintf("error scanning row: %v\n", err))
		}
		*dest = append(*dest, u)
	}

	return nil
}

func (u *User) Get(username string) error {
	var (
		err error
		row *sql.Row
	)
	sqlStatement := `SELECT id,username,password
					FROM users
					WHERE username=$1
					`
	row = u.conn.QueryRow(sqlStatement, username)
	switch err = row.Scan(&u.Id, &u.Username, &u.Password); err {
	case sql.ErrNoRows:
		return errors.New("no row where retrieved")
	case nil:
		return nil
	default:
		return errors.New(fmt.Sprintf("error retrieving user from database: %v\n", err))
	}
}
