package models

type Transaction struct {
	Base
	Customer Customer `json:"customer"`
	CustomerID ID `gorm:"not null" json:"customer_id"`
	Description string `gorm:"size:255" json:"description"`
}
