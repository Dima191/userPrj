package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"ID"`
	Name         string `json:"Name"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	HashPassword string `json:"HashPassword"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Length(2, 15)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(u.validatePassword(u.HashPassword == "")), validation.Length(8, 100)),
	)

}

func (u *User) validatePassword(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

func (u *User) BeforeCreate() error {
	if len(u.Password) != 0 {
		hashPassword, err := EncryptPassword(u.Password)
		if err != nil {
			return err
		}
		u.HashPassword = hashPassword
		u.Password = ""
	} else {

	}
	return nil
}

func EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ComparePasswords(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
