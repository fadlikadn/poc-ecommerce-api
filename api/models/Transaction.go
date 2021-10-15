package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Transaction struct {
	Base
	Customer    Customer `json:"customer"`
	CustomerID  ID       `gorm:"not null" json:"customer_id"`
	Description string   `gorm:"size:255" json:"description"`
}

func (t *Transaction) Prepare() {
	t.ID = 0
	t.Customer = Customer{}
	t.Description = html.EscapeString(strings.TrimSpace(t.Description))
}

func (t *Transaction) Validate() error {
	if t.CustomerID < 1 {
		return errors.New("Required Customer ID")
	}
	return nil
}

func (t *Transaction) AddNewTransaction(db *gorm.DB) (*Transaction, error) {
	var err error
	err = db.Debug().Model(&Transaction{}).Create(&t).Error
	if err != nil {
		return &Transaction{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&Customer{}).Where("id = ?", t.CustomerID).Take(&t.Customer).Error
		if err != nil {
			return &Transaction{}, err
		}
	}
	return t, nil
}

func (t *Transaction) FindAllTransactionsByCustomer(db *gorm.DB, cid uint64) (*[]Transaction, error) {
	var err error
	var transactions []Transaction
	err = db.Debug().Model(&Transaction{}).Where("customer_id = ?", cid).Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	if len(transactions) > 0 {
		for i, _ := range transactions {
			err := db.Debug().Model(&Customer{}).Where("id = ?", cid).Take(&transactions[i].Customer).Error
			if err != nil {
				return &[]Transaction{}, err
			}
		}
	}
	return &transactions, nil
}
