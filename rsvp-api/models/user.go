package models

import (
	"errors"

	"github.com/temiloluwa/rsvp-api/db"
	hash "github.com/temiloluwa/rsvp-api/utils"
)

type User struct {
	ID       int
	Name     string
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
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	u.Password, _ = hash.HashPassword(u.Password)
	result, err := db.DB.Exec(query, u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = int(userId)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	var retrievedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}
	isValid := hash.CheckPasswordHash(retrievedPassword, u.GetPassword())

	if !isValid {
		return errors.New("invalid credentials")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT id, email, password FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
