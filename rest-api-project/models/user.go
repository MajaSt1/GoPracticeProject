package models

import (
	"errors"

	"example.com/note/rest-api-project/db"
	"example.com/note/rest-api-project/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPasswored, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPasswored)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err // nil / error
}

func (u User) ValidateCredentials() error {
	query := "SELECT email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email) // we know that we query exactly one row
	var retrievedPwd string
	err := row.Scan(&retrievedPwd)
	if err != nil {
		return errors.New("Credentials invalid.")
	}
	
	passwotdIsValid := utils.CheckPasswordHash(u.Password, retrievedPwd)
	if !passwotdIsValid {
		return errors.New("Credentials invalid.")
	}

	return nil
}