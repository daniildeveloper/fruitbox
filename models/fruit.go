package models

import (
	"github.com/jinzhu/gorm"
)

type Fruit struct {
	gorm.Model
	Name         string `gorm:"not null;unique" json:"name"`
	PricePerUnit float32
	UnitID       int
	Unit         FruitUnit
}

type FruitUnit struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}
