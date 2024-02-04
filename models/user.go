package models

import (
	"errors"
	"github.com/Tijanieneye10/db"
	"github.com/Tijanieneye10/utils"
)

type User struct {
	ID       int
	Name     string
	Password string `binding:"required"`
	Email    string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(name, password, email) VALUES(
		  ?,?,?                                             
		)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, hashPassword, u.Email)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = int(userId)
	u.Password = hashPassword

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isValidPassword := utils.CheckPassword(u.Password, retrievedPassword)

	if !isValidPassword {
		return errors.New("invalid credentials")
	}

	return nil
}
