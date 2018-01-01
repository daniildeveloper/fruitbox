package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Client -
type Client struct {
	gorm.Model
	Name  string
	Phone string
	Email string
}

// Address -
type Address struct {
	gorm.Model
	Client   Client `gorm:"ForeignKey:ClientID"`
	ClientID int
	Address  string //TODO: here type text
	Notes    string `gorm:"type:text"`
}

// Proposl - заявка
type Proposal struct {
	gorm.Model
	Client             Client
	ClientID           int
	Address            Address
	AddressString      string
	ProposalJSONString string
	Deadline           time.Time
}

// ProposalManager -
type ProposalManager struct {
	db *DB
}

// NewProposalManager -
func NewProposalManager(db *DB) (*ProposalManager, error) {
	manager := ProposalManager{}
	manager.db = db

	db.AutoMigrate(
		&Client{},
		&Address{},
		&Proposal{})

	return &manager, nil
}
