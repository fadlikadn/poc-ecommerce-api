package models

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type WarehouseItem struct {
	Base
	SupplierItem   SupplierItem `json:"supplier_item"`
	SupplierItemID ID           `gorm:"not null" json:"supplier_item_id"`
	SKU            string       `gorm:"size:255;not null;unique" json:"sku"`
	BrandName      string       `gorm:"size:100;not null" json:"brand_name"`
	ModelName      string       `gorm:"size:100;not null;unique" json:"model_name"`
	Description    string       `gorm:"size:255" json:"description"`
	Price          IDR          `gorm:"not null" json:"price"`
	Quantity       Amount       `gorm:"not null" json:"quantity"`
}

func (w *WarehouseItem) Prepare() {
	w.ID = 0
	w.SKU = html.EscapeString(strings.TrimSpace(w.SKU))
	w.BrandName = html.EscapeString(strings.TrimSpace(w.BrandName))
	w.ModelName = html.EscapeString(strings.TrimSpace(w.ModelName))
	w.Description = html.EscapeString(strings.TrimSpace(w.Description))
}

func (w *WarehouseItem) FindWarehouseByID(db *gorm.DB, wid ID) (*WarehouseItem, error) {
	var err error
	err = db.Debug().Model(&WarehouseItem{}).Where("id = ?", wid).Take(&w).Error
	if err != nil {
		return &WarehouseItem{}, err
	}
	if w.ID != 0 {
		err = db.Debug().Model(&SupplierItem{}).Where("id = ?", w.SupplierItemID).Take(&w.SupplierItem).Error
		if err != nil {
			return &WarehouseItem{}, err
		}
	}
	return w, nil
}
