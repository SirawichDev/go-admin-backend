package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           string `json:"id"`
	Firstname    string `json:"first_name"`
	Lastname     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	IsAmbassador bool   `json:"-"`
	Password     []byte `json:"-"`
}

func (user *User) SetPassword(password string) {
	cost := 12
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
