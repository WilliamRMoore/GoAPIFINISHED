package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email,password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedpw, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedpw)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPw string
	err := row.Scan(&u.ID, &retrievedPw)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	pwIsValid := utils.CheckPasswordHash(u.Password, retrievedPw)

	if !pwIsValid {
		return errors.New("Invalid credentials")
	}

	return nil

}
