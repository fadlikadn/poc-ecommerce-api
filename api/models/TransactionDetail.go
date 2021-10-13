package models

type TransactionDetail struct {
	Base
	Transaction Transaction `json:"transaction"`
	TransactionID ID `gorm:"not null" json:"transaction_id"`
	WarehouseItem WarehouseItem `json:"warehouseItem"`
	WarehouseItemID ID `gorm:"not null" json:"warehouse_item_id"`
	Price IDR `gorm:"not null" json:"price"`
	Quantity Amount `gorm:"not null" json:"quantity"`
}
