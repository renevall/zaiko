package domain

import (
	"github.com/jinzhu/gorm"
)

// Trans model definition
type Trans struct {
	gorm.Model
	ProductID int    `json:"product_id"`
	Amount    int    `json:"amount"`
	Action    string `json:"action"`
}

//TransStore is the contract interface to the DB
type TransStore interface {
	CreateTrans(t *Trans) error
	ReadTransByID(id int) Trans
	ReadTransByProduct(id int) []Trans
	UpdateTrans(t Trans) Trans
}
