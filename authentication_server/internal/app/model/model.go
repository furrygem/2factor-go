// package model provides models for different db tables
package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username      string `json:"username"`
	ClearPassword string `json:"clear_password"`
	Password      string `json:"password"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.ClearPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return nil
}
