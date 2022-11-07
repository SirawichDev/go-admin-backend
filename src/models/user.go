package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        string
	Firstname string
	Lastname  string
	Email     string
	IsAdmin   bool
	Password  []byte
}

func (user *User) SetPassword(password string) {
	cost := 12
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
