package models

import (
	"github.com/jinzhu/gorm"
)

// Fruit - базвоая единица. Компонуем, управляем
type Fruit struct {
	gorm.Model
	Name         string `gorm:"not null;unique" json:"name"`
	PricePerUnit float32
	UnitID       int
	Unit         FruitUnit
}

// FruitUnit - Единица измерения фрукта. КГ, шт, г,
type FruitUnit struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}

// FruitManager -
type FruitManager struct {
	db *DB
}

// NewFruitManager -
func NewFruitManager(db *DB) (*FruitManager, error) {
	manager := FruitManager{}
	manager.db = db

	db.AutoMigrate(
		&Fruit{},
		&FruitUnit{})

	return &manager, nil
}
