package models

import (
	"github.com/google/uuid"
	"time"
)

type IDR int64
type Amount int
type UUID uuid.UUID
type ID uint64

// Base contains common columns for all tables.
type Base struct {
	//ID UUID `gorm:"type:uuid;primary_key;"`
	ID        ID `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID
//func (base *Base) BeforeCreate(scope *gorm.Scope) error {
//	uuid, err := uuid.NewRandom()
//	if err != nil {
//		return err
//	}
//	return scope.SetColumn("ID", uuid)
//}
