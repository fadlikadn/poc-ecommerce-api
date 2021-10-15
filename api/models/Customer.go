package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type Customer struct {
	Base
	Nickname string `gorm:"size:255;not null;unique" json:"nickname"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;unique" json:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (c *Customer) BeforeSave() error {
	hashedPassword, err := Hash(c.Password)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)
	return nil
}

func (c *Customer) Prepare() {
	c.ID = 0
	c.Nickname = html.EscapeString(strings.TrimSpace(c.Nickname))
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
}

func (c *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if c.Password == "" {
			return errors.New("Required Password")
		}
		if c.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		return nil
	}
}
