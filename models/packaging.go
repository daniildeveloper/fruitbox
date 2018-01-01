package models

import "github.com/jinzhu/gorm"

// Packaging - упаковка
type Packaging struct {
	gorm.Model
	Name   string
	Price  float32
	Weight float32
	Width  float32
	Height float32
}

// PackagingFruits -
type PackagingFruits struct {
	gorm.Model
	Fruit       Fruit
	FruitID     int
	Packaging   Packaging
	PackagingID int
	Units       float32
}

// PackagingManager -
type PackagingManager struct {
	db *DB
}

// NewPackagingManager -
func NewPackagingManager(db *DB) (*PackagingManager, error) {
	manager := PackagingManager{}
	manager.db = db

	db.AutoMigrate(
		&Packaging{},
		&PackagingFruits{})

	return &manager, nil
}
