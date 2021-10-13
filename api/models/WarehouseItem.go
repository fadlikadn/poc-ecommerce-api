package models

type WarehouseItem struct {
	Base
	SupplierItem SupplierItem `json:"supplier_item"`
	SupplierItemID ID `gorm:"not null" json:"supplier_item_id"`
	SKU string `gorm:"size:255;not null;unique" json:"sku"`
	BrandName string `gorm:"size:100;not null;unique" json:"brand_name"`
	ModelName string `gorm:"size:100;not null;unique" json:"model_name"`
	Description string `gorm:"size:255" json:"description"`
	Price IDR `gorm:"not null" json:"price"`
	Quantity Amount `gorm:"not null" json:"quantity"`
}
