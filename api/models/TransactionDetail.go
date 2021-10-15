package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type TransactionDetail struct {
	Base
	Transaction     Transaction   `json:"transaction"`
	TransactionID   ID            `gorm:"not null" json:"transaction_id"`
	WarehouseItem   WarehouseItem `json:"warehouseItem"`
	WarehouseItemID ID            `gorm:"not null" json:"warehouse_item_id"`
	Price           IDR           `gorm:"not null" json:"price"`
	Quantity        Amount        `gorm:"not null" json:"quantity"`
}

func (t *TransactionDetail) Prepare() {
	t.ID = 0
	t.Transaction = Transaction{}
	t.WarehouseItem = WarehouseItem{}
}

func (t *TransactionDetail) Validate(action string) error {
	switch strings.ToLower(action) {
	case "toCart":
		// validate to cart process
		if t.TransactionID < 1 {
			return errors.New("Required Transaction ID")
		}
		if t.WarehouseItemID < 1 {
			return errors.New("Required WarehouseItem ID")
		}
		return nil
	case "checkout":
		// validate checkout process
		// check current stock by count all purchased warehouseItem
		return nil
	default:
		return nil
	}
}

func (t *TransactionDetail) AddNewTransactionDetail(db *gorm.DB) (*TransactionDetail, error) {
	var err error
	err = db.Debug().Model(&TransactionDetail{}).Create(&t).Error
	if err != nil {
		return &TransactionDetail{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&Transaction{}).Where("id = ?", t.TransactionID).Take(&t.Transaction).Error
		if err != nil {
			return &TransactionDetail{}, err
		}
	}
	return t, nil
}

func (t *TransactionDetail) UpdateTransactionDetail(db *gorm.DB) (*TransactionDetail, error) {
	var err error
	err = db.Debug().Model(&TransactionDetail{}).Where("id = ?", t.ID).Updates(TransactionDetail{Quantity: t.Quantity, Price: t.Price}).Error
	if err != nil {
		return &TransactionDetail{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&Transaction{}).Where("id = ?", t.TransactionID).Take(&t.Transaction).Error
		if err != nil {
			return &TransactionDetail{}, err
		}
		err = db.Debug().Model(&WarehouseItem{}).Where("id = ?", t.WarehouseItemID).Take(&t.WarehouseItem).Error
		if err != nil {
			return &TransactionDetail{}, err
		}
	}
	return t, nil
}

func (t *TransactionDetail) DeleteTransactionDetail(db *gorm.DB, tid ID, wid ID) (int64, error) {
	db = db.Debug().Model(&TransactionDetail{}).Where("transaction_id = ? and warehouse_item_id = ?", tid, wid).Take(&TransactionDetail{}).Delete(&TransactionDetail{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Transaction Detail not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (t *TransactionDetail) FindByTransaction(db *gorm.DB, tid ID) (*[]TransactionDetail, error) {
	var err error
	var transactions []TransactionDetail
	err = db.Debug().Model(&TransactionDetail{}).Where("transaction_id = ?", tid).Find(&transactions).Error
	if err != nil {
		return &[]TransactionDetail{}, err
	}
	return &transactions, err
}

func (t *TransactionDetail) GetWarehouseItemTotalPurchased(db *gorm.DB, wid ID) (Amount, error) {
	var err error
	var transactions []TransactionDetail
	var total Amount = 0
	err = db.Debug().Model(&TransactionDetail{}).Where("warehouse_item_id = ?", wid).Find(&transactions).Error
	if err != nil {
		return 0, err
	}
	if len(transactions) > 0 {
		for i, _ := range transactions {
			total = total + transactions[i].Quantity
		}
	}
	return total, nil
}
