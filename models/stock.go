package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Inventory -
type Inventory struct {
	gorm.Model
	User      User
	UserID    int
	DateBegin time.Time
	DateEnd   time.Time
}

// FruitInventory - проверка остатков по каждому фрукту
type FruitInventory struct {
	gorm.Model
	Inventory     Inventory
	InventoryID   int
	Fruit         Fruit
	FruitID       int
	ExpectedCount float32
	RealCount     float32
	Diffirence    float32
}

// Wiki - wikipedia like markdown for each fruit
type Wiki struct {
	gorm.Model
	Fruit   Fruit
	FruitID int
	Text    string `gorm:"type:text"`
}

// StockManager -
type StockManager struct {
	db *DB
}

// NewStockManager -
func NewStockManager(db *DB) (*StockManager, error) {
	manager := StockManager{}
	manager.db = db

	db.AutoMigrate(
		&Inventory{},
		&FruitInventory{},
		&Wiki{})

	return &manager, nil
}
