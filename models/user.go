package models

import (
	"errors"
	"net/mail"
	"time"
)

type User struct {
	Id             int64     `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Deleted        bool      `json:"deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *User) ValidateFields() error {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("email inválido")
	}
	if u.FirstName == "" {
		return errors.New("nombre inválido")
	}
	if u.LastName == "" {
		return errors.New("apellido inválido")
	}
	if u.HashedPassword == "" {
		return errors.New("contraseña inválida")
	}

	return nil
}
