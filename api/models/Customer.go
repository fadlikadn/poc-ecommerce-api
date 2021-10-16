package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"strings"
	"time"
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
	case "update":
		if c.Nickname == ""	{
			return errors.New("Required Nickname")
		}
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
		if c.Nickname == "" {
			return errors.New("Required Nickname")
		}
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
	}
}

func (c *Customer) AddCustomer(db *gorm.DB) (*Customer, error) {
	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

func (c *Customer) FindAllCustomers(db *gorm.DB) (*[]Customer, error) {
	var err error
	customers := []Customer{}
	err = db.Debug().Model(&Customer{}).Limit(100).Find(&customers).Error
	if err != nil {
		return &[]Customer{}, err
	}
	return &customers, err
}

func (c *Customer) FindCustomerByID(db *gorm.DB, cid ID) (*Customer, error) {
	var err error
	err = db.Debug().Model(&Customer{}).Where("id = ?", cid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer Not Found")
	}
	return c, err
}

func (c *Customer) UpdateCustomer(db *gorm.DB, cid ID) (*Customer, error) {
	// To hash the password
	err := c.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Customer{}).Where("id = ?", cid).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"password": c.Password,
			"nickname": c.Nickname,
			"email": c.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Customer{}, db.Error
	}

	// This is the display the updated customer
	err = db.Debug().Model(&Customer{}).Where("id = ?", cid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}


func (c *Customer) DeleteCustomer(db *gorm.DB, cid ID) (int64, error) {
	db = db.Debug().Model(&Customer{}).Where("id = ?", cid).Take(&Customer{}).Delete(&Customer{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}