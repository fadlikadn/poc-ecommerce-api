package models

type ObjectSchema struct {
	Base
	WarehouseItem   WarehouseItem `json:"warehouse_item"`
	WarehouseItemID ID            `gorm:"not null" json:"warehouse_item_id"`
	Description     string        `gorm:"size:255" json:"description"`
	Type            string        `gorm:"size:100;not null" json:"type"`
	Value           string        `gorm:"size:100;not null" json:"value"`
}
