package models

import (
	"github.com/temiloluwa/1-finance-tracker-api/db"
	hash "github.com/temiloluwa/1-finance-tracker-api/utils"
)

type User struct {
	ID       int
	Password string `binding:"required"`
	Email    string `binding:"required"`
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	u.Password, _ = hash.HashPassword(u.Password)
	result, err := db.DB.Exec(query, u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = int(userId)
	return err
}
