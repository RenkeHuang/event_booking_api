package models

import (
	"errors"
	"event_booking/db"
	"event_booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	password string `binding:"required"`
}

func (user *User) Save() error {
	// Store the user to the database
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the user data
	hashedPassword, err := utils.HashPassword(user.password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	user.ID = int64(id)

	return err
}

func (u *User) ValidateCredentials() error {
	// Check if the user exists in the database
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return errors.New("Invalid credentials" + err.Error())
	}

	// Compare the provided password with the hashed password
	passwordIsValid := utils.CheckPasswordHash(hashedPassword, u.password)
	if !passwordIsValid {
		return errors.New("Invalid credentials" + hashedPassword + " " + u.password)
	}

	return nil
}
