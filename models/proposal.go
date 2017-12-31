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
}

type Proposal struct {
	gorm.Model
	Client             Client
	ClientID           int
	Address            Address
	AddressString      string
	ProposalJSONString string
	Deadline           time.Time
}
